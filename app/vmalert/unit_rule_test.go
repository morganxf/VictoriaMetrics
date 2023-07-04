package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"syscall"
	"testing"
	"time"

	"gopkg.in/yaml.v2"

	"github.com/VictoriaMetrics/VictoriaMetrics/lib/fs"
	// "github.com/VictoriaMetrics/metricsql"

	testutil "github.com/VictoriaMetrics/VictoriaMetrics/app/victoria-metrics/test"
	"github.com/VictoriaMetrics/VictoriaMetrics/app/vmalert/datasource"
	"github.com/VictoriaMetrics/VictoriaMetrics/app/vmalert/remotewrite"
	"github.com/VictoriaMetrics/VictoriaMetrics/app/vminsert"
	"github.com/VictoriaMetrics/VictoriaMetrics/app/vmselect"
	"github.com/VictoriaMetrics/VictoriaMetrics/app/vmselect/promql"
	"github.com/VictoriaMetrics/VictoriaMetrics/app/vmstorage"

	// "github.com/VictoriaMetrics/VictoriaMetrics/lib/fs"
	// "github.com/VictoriaMetrics/VictoriaMetrics/lib/httpserver"
	vmalertconfig "github.com/VictoriaMetrics/VictoriaMetrics/app/vmalert/config"
	"github.com/VictoriaMetrics/VictoriaMetrics/app/vmalert/notifier"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/httpserver"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/logger"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/prompbmarshal"

	// "github.com/VictoriaMetrics/metricsql"

	prommodel "github.com/prometheus/common/model"
	promlabel "github.com/prometheus/prometheus/model/labels"
	promparser "github.com/prometheus/prometheus/promql/parser"
)

func TestUnitRule(t *testing.T) {
	setUp()
	RulesUnitTest([]string{"./ruletest/test.yaml"}...)
	tearDown()
}

var testStartTime = time.Unix(0, 0).UTC()

var (
	storagePath   string
	insertionTime = time.Now().UTC()
)

const (
	testFixturesDir            = "ruletest"
	testStorageSuffix          = "vm-test-storage"
	testHTTPListenAddr         = ":7655"
	testStatsDListenAddr       = ":2004"
	testOpenTSDBListenAddr     = ":4244"
	testOpenTSDBHTTPListenAddr = ":4245"
	testLogLevel               = "INFO"
)

const (
	testReadHTTPPath          = "http://127.0.0.1" + testHTTPListenAddr
	testWriteHTTPPath         = "http://127.0.0.1" + testHTTPListenAddr + "/write"
	testOpenTSDBWriteHTTPPath = "http://127.0.0.1" + testOpenTSDBHTTPListenAddr + "/api/put"
	testPromWriteHTTPPath     = "http://127.0.0.1" + testHTTPListenAddr + "/api/v1/write"
	testHealthHTTPPath        = "http://127.0.0.1" + testHTTPListenAddr + "/health"
)

func processFlags() {
	flag.Parse()
	for _, fv := range []struct {
		flag  string
		value string
	}{
		{flag: "storageDataPath", value: storagePath},
		{flag: "httpListenAddr", value: testHTTPListenAddr},
		{flag: "graphiteListenAddr", value: testStatsDListenAddr},
		{flag: "opentsdbListenAddr", value: testOpenTSDBListenAddr},
		{flag: "loggerLevel", value: testLogLevel},
		{flag: "opentsdbHTTPListenAddr", value: testOpenTSDBHTTPListenAddr},
		{flag: "storageDataPath", value: storagePath},
		// set storage retention time to 100 years
		{flag: "retentionPeriod", value: "1200"},
		{flag: "datasource.url", value: "http://127.0.0.1:7655/prometheus"},
		{flag: "remoteWrite.url", value: "http://127.0.0.1:7655"},
	} {
		// panics if flag doesn't exist
		if err := flag.Lookup(fv.flag).Value.Set(fv.value); err != nil {
			log.Fatalf("unable to set %q with value %q, err: %v", fv.flag, fv.value, err)
		}
	}
}

func testrequestHandler(w http.ResponseWriter, r *http.Request) bool {
	if r.URL.Path == "/" {
		if r.Method != http.MethodGet {
			return false
		}
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, "<h2>Single-node VictoriaMetrics</h2></br>")
		fmt.Fprintf(w, "See docs at <a href='https://docs.victoriametrics.com/'>https://docs.victoriametrics.com/</a></br>")
		fmt.Fprintf(w, "Useful endpoints:</br>")
		httpserver.WriteAPIHelp(w, [][2]string{
			{"vmui", "Web UI"},
			{"targets", "status for discovered active targets"},
			{"service-discovery", "labels before and after relabeling for discovered targets"},
			{"metric-relabel-debug", "debug metric relabeling"},
			{"expand-with-exprs", "WITH expressions' tutorial"},
			{"api/v1/targets", "advanced information about discovered targets in JSON format"},
			{"config", "-promscrape.config contents"},
			{"metrics", "available service metrics"},
			{"flags", "command-line flags"},
			{"api/v1/status/tsdb", "tsdb status page"},
			{"api/v1/status/top_queries", "top queries"},
			{"api/v1/status/active_queries", "active queries"},
		})
		return true
	}
	if vminsert.RequestHandler(w, r) {
		return true
	}
	if vmselect.RequestHandler(w, r) {
		return true
	}
	if vmstorage.RequestHandler(w, r) {
		logger.Infof("wang %s", "test")
		return true
	}
	return false
}

func waitFor(timeout time.Duration, f func() bool) error {
	fraction := timeout / 10
	for i := fraction; i < timeout; i += fraction {
		if f() {
			return nil
		}
		time.Sleep(fraction)
	}
	return fmt.Errorf("timeout")
}

const (
	testStorageInitTimeout = 10 * time.Second
)

func setUp() {
	storagePath = filepath.Join(os.TempDir(), testStorageSuffix)
	processFlags()
	logger.Init()
	vmstorage.Init(promql.ResetRollupResultCacheIfNeeded)
	vmselect.Init()
	vminsert.Init()
	go httpserver.Serve(*httpListenAddr, false, testrequestHandler)
	readyStorageCheckFunc := func() bool {
		resp, err := http.Get(testHealthHTTPPath)
		if err != nil {
			return false
		}
		_ = resp.Body.Close()
		return resp.StatusCode == 200
	}
	if err := waitFor(testStorageInitTimeout, readyStorageCheckFunc); err != nil {
		log.Fatalf("http server can't start for %s seconds, err %s", testStorageInitTimeout, err)
	}
}

func tearDown() {
	if err := httpserver.Stop(*httpListenAddr); err != nil {
		log.Printf("cannot stop the webservice: %s", err)
	}
	vminsert.Stop()
	vmstorage.Stop()
	vmselect.Stop()
	fs.MustRemoveAll(storagePath)
}

// RulesUnitTest does unit testing of rules based on the unit testing files provided.
// More info about the file format can be found in the docs.
func RulesUnitTest(files ...string) int {
	failed := false

	for _, f := range files {
		if errs := ruleUnitTest(f); errs != nil {
			fmt.Fprintln(os.Stderr, "  FAILED:")
			for _, e := range errs {
				fmt.Fprintln(os.Stderr, e.Error())
			}
			failed = true
		} else {
			fmt.Println("  SUCCESS")
		}
	}
	if failed {
		return 1
	}
	return 0
}

// resolveAndGlobFilepaths joins all relative paths in a configuration
// with a given base directory and replaces all globs with matching files.
func resolveAndGlobFilepaths(baseDir string, utf *unitTestFile) error {
	for i, rf := range utf.RuleFiles {
		if rf != "" && !filepath.IsAbs(rf) {
			utf.RuleFiles[i] = filepath.Join(baseDir, rf)
		}
	}

	var globbedFiles []string
	for _, rf := range utf.RuleFiles {
		m, err := filepath.Glob(rf)
		if err != nil {
			return err
		}
		if len(m) == 0 {
			fmt.Fprintln(os.Stderr, "  WARNING: no file match pattern", rf)
		}
		globbedFiles = append(globbedFiles, m...)
	}
	utf.RuleFiles = globbedFiles
	return nil
}

func ruleUnitTest(filename string) []error {
	fmt.Println("Unit Testing: ", filename)

	b, err := os.ReadFile(filename)
	if err != nil {
		return []error{err}
	}

	var unitTestInp unitTestFile
	if err := yaml.UnmarshalStrict(b, &unitTestInp); err != nil {
		return []error{err}
	}
	if err := resolveAndGlobFilepaths(filepath.Dir(filename), &unitTestInp); err != nil {
		return []error{err}
	}

	if unitTestInp.EvaluationInterval == 0 {
		unitTestInp.EvaluationInterval = prommodel.Duration(1 * time.Minute)
	}

	evalInterval := time.Duration(unitTestInp.EvaluationInterval)

	// Giving number for groups mentioned in the file for ordering.
	// Lower number group should be evaluated before higher number group.
	groupOrderMap := make(map[string]int)
	for i, gn := range unitTestInp.GroupEvalOrder {
		if _, ok := groupOrderMap[gn]; ok {
			return []error{fmt.Errorf("group name repeated in evaluation order: %s", gn)}
		}
		groupOrderMap[gn] = i
	}

	// Testing.
	var errs []error
	for _, t := range unitTestInp.Tests {
		ers := t.test(evalInterval, groupOrderMap, unitTestInp.RuleFiles...)
		if ers != nil {
			errs = append(errs, ers...)
		}
	}

	if len(errs) > 0 {
		return errs
	}
	return nil
}

// unitTestFile holds the contents of a single unit test file.
type unitTestFile struct {
	RuleFiles          []string           `yaml:"rule_files"`
	EvaluationInterval prommodel.Duration `yaml:"evaluation_interval,omitempty"`
	GroupEvalOrder     []string           `yaml:"group_eval_order"`
	Tests              []testGroup        `yaml:"tests"`
}

// testGroup is a group of input series and tests associated with it.
type testGroup struct {
	Interval        prommodel.Duration `yaml:"interval"`
	InputSeries     []series           `yaml:"input_series"`
	AlertRuleTests  []alertTestCase    `yaml:"alert_rule_test,omitempty"`
	PromqlExprTests []promqlTestCase   `yaml:"promql_expr_test,omitempty"`
	ExternalLabels  promlabel.Labels   `yaml:"external_labels,omitempty"`
	ExternalURL     string             `yaml:"external_url,omitempty"`
	TestGroupName   string             `yaml:"name,omitempty"`
}

// maxEvalTime returns the max eval time among all alert and promql unit tests.
func (tg *testGroup) maxEvalTime() time.Duration {
	var maxd prommodel.Duration
	for _, alert := range tg.AlertRuleTests {
		if alert.EvalTime > maxd {
			maxd = alert.EvalTime
		}
	}
	for _, pet := range tg.PromqlExprTests {
		if pet.EvalTime > maxd {
			maxd = pet.EvalTime
		}
	}
	return time.Duration(maxd)
}

type series struct {
	Series string `yaml:"series"`
	Values string `yaml:"values"`
}

type alertTestCase struct {
	EvalTime  prommodel.Duration `yaml:"eval_time"`
	Alertname string             `yaml:"alertname"`
	ExpAlerts []alert            `yaml:"exp_alerts"`
}

type alert struct {
	ExpLabels      map[string]string `yaml:"exp_labels"`
	ExpAnnotations map[string]string `yaml:"exp_annotations"`
}

type promqlTestCase struct {
	Expr       string             `yaml:"expr"`
	EvalTime   prommodel.Duration `yaml:"eval_time"`
	ExpSamples []sample           `yaml:"exp_samples"`
}

type sample struct {
	Labels string  `yaml:"labels"`
	Value  float64 `yaml:"value"`
}

func httpWrite(address string, r io.Reader) {
	resp, err := http.Post(address, "", r)
	if err != nil || resp.StatusCode != 204 {
		logger.Errorf("failed to send to storage: %v", err)
	}
	resp.Body.Close()
}

func (tg *testGroup) test(evalInterval time.Duration, groupOrderMap map[string]int, ruleFiles ...string) []error {
	// todo defer cleanup data
	r := testutil.WriteRequest{}
	for _, data := range tg.InputSeries {
		result := fmt.Sprintf("%v %v\n", data.Series, data.Values)
		prommetric, promvals, err := promparser.ParseSeriesDesc(result)
		if err != nil {
			logger.Errorf("failed to parse series %v", err)
		}
		// fmt.Println(prommetric, promvals)
		// expr, err := metricsql.Parse(data.Series)
		// fmt.Println(expr)
		// if err != nil {
		// }
		// exp := expr.(*metricsql.MetricExpr)
		samples := make([]testutil.Sample, 0, len(promvals))
		ts := testStartTime
		for _, v := range promvals {
			if !v.Omitted {
				samples = append(samples, testutil.Sample{
					Timestamp: ts.UnixNano() / int64(time.Millisecond/time.Nanosecond),
					Value:     v.Value,
				})
			}
			ts = ts.Add(time.Duration(tg.Interval))
		}
		var ls []testutil.Label
		for _, filter := range prommetric {
			ls = append(ls, testutil.Label{Name: filter.Name, Value: filter.Value})
		}
		r.Timeseries = append(r.Timeseries, testutil.TimeSeries{Labels: ls, Samples: samples})
	}
	data, err := testutil.Compress(r)
	if err != nil {
		logger.Errorf("error compressing %v %s", r, err)
	}
	httpWrite(testPromWriteHTTPPath, bytes.NewBuffer(data))
	// flush replay result
	vmstorage.Storage.DebugFlush()

	testGroups, err := vmalertconfig.Parse(ruleFiles, notifier.ValidateTemplates, true)
	if err != nil {
		return []error{err}
	}

	var alertEvalTimes []prommodel.Duration
	alertResultMap := map[prommodel.Duration]map[string][]alert{}
	for _, at := range tg.AlertRuleTests {
		alertEvalTimes = append(alertEvalTimes, prommodel.Duration(at.EvalTime))
		if _, ok := alertResultMap[at.EvalTime]; !ok {
			alertResultMap[at.EvalTime] = make(map[string][]alert)
		}
		alertResultMap[at.EvalTime][at.Alertname] = at.ExpAlerts
	}
	sort.Slice(alertEvalTimes, func(i, j int) bool {
		return alertEvalTimes[i] < alertEvalTimes[j]
	})

	mint := time.Unix(0, 0)
	maxt := mint.Add(tg.maxEvalTime())

	q, err := datasource.Init(url.Values{"nocache": {"1"}})
	if err != nil {
		return []error{err}
	}
	rw, err := remotewrite.Init(context.Background())
	if err != nil {
		logger.Fatalf("failed to init remoteWrite: %s", err)
	}

	logger.Infof("will sleep here %s", "wang")
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// wait here so we can using vmui
	for {
		select {
		case s := <-sigs:
			logger.Infof("program will exit now cause receiving signal %s", s)
			os.Exit(1)
		default:
		}
	}

	var groups []*Group
	for _, g := range testGroups {
		ng := newGroup(g, q, *evaluationInterval, nil)
		groups = append(groups, ng)
		// should got alert for rule InstanceUp/InstanceLongUp
		// todo 弃用replay，跟prom一样先用query跑通，后面测试replay的功能[replay目前跑出来的数据不对]

		// num := ng.replay(mint, maxt, rw)
		// logger.Infof("replay got %d results", num)
	}
	e := &executor{
		rw: rw,
		notifiers: func() []notifier.Notifier {
			return []notifier.Notifier{}
		},
		notifierHeaders:          make(map[string]string),
		previouslySentSeriesToRW: make(map[uint64]map[string][]prompbmarshal.Label),
	}
	var checkErrs []error
	curr := 0
	for ts := mint; ts.Before(maxt) || ts.Equal(maxt); ts = ts.Add(evalInterval) {
		for _, g := range groups {
			resolveDuration := getResolveDuration(g.Interval, *resendDelay, *maxResolveDuration)
			// g.Rules.alert.Restored
			errs := e.execConcurrently(context.Background(), g.Rules, ts, g.Concurrency, resolveDuration, g.Limit)
			for err := range errs {
				if err != nil {
					checkErrs = append(checkErrs, fmt.Errorf("    group: %q, time: %s, err: %w", g.Name,
						ts, err))
				}
			}
		}
		// getRealAlert := func(name string) *AlertingRule {
		// 	for _, grule := range g.Rules {
		// 		if alertrule, ok := grule.(*AlertingRule); ok {
		// 			if alertrule.Name == name {
		// 				return alertrule
		// 			}
		// 		}
		// 	}
		// 	return nil
		// }
		// 	如果时间点需要检查，检查之前flush？
		// 检查告警的时候什么时候需要flush，每一次时间循环吗
		// vmstorage.Storage.DebugFlush()

		for curr < len(alertEvalTimes) {
			if ts.Sub(mint) > time.Duration(alertEvalTimes[curr]) ||
				time.Duration(alertEvalTimes[curr]) >= ts.Add(evalInterval).Sub(mint) {
				break
			}
			curr++

			needToCheck := map[string]labelsAndAnnotations{}
			for _, g := range groups {
				for _, grule := range g.Rules {
					if alertrule, ok := grule.(*AlertingRule); ok {
						if _, ok := alertResultMap[prommodel.Duration(ts.UnixNano())][alertrule.Name]; ok {
							for _, got := range alertrule.alerts {
								if got.State != notifier.StateFiring {
									continue
								}
								delete(got.Labels, "alertgroup")
								needToCheck[alertrule.Name] = append(needToCheck[alertrule.Name], labelAndAnnotation{
									Labels:      convertToLabels(got.Labels),
									Annotations: convertToLabels(got.Annotations),
								})
							}
						}
					}
				}
			}

			for alertName, res := range alertResultMap[alertEvalTimes[curr]] {
				// realAlert := getRealAlert(alertName)
				// if len(exp) == 0 {
				// 	if realAlert == nil {
				// 		continue
				// 	}
				// 	// todo append error
				// }
				// notifyAlerts := realAlert.alertsToSend(ts, resolveDuration, *resendDelay)
				// todo prometheus是要求全部匹配还是部分匹配就好了
				gotAlerts := needToCheck[alertName]
				var expAlerts labelsAndAnnotations
				for _, expAlert := range res {
					// User gives only the labels from alerting rule, which doesn't
					// include this label (added by Prometheus during Eval).
					if expAlert.ExpLabels == nil {
						expAlert.ExpLabels = make(map[string]string)
					}
					expAlert.ExpLabels[alertNameLabel] = alertName
					expAlerts = append(expAlerts, labelAndAnnotation{
						Labels:      convertToLabels(expAlert.ExpLabels),
						Annotations: convertToLabels(expAlert.ExpAnnotations),
					})
				}
				sort.Sort(gotAlerts)
				sort.Sort(expAlerts)
				if !reflect.DeepEqual(expAlerts, gotAlerts) {
					var testName string
					if tg.TestGroupName != "" {
						testName = fmt.Sprintf("    name: %s,\n", tg.TestGroupName)
					}
					expString := indentLines(expAlerts.String(), "            ")
					gotString := indentLines(gotAlerts.String(), "            ")
					checkErrs = append(checkErrs, fmt.Errorf("%s    alertname: %s, time: %s, \n        exp:%v, \n        got:%v",
						testName, alertName, alertEvalTimes[curr].String(), expString, gotString))
				}
			}
			// 在需要检查的时间戳下，检查需要检查的告警结果
		}

	}
	// 需要吗
	vmstorage.Storage.DebugFlush()

	queries := q.BuildWithParams(datasource.QuerierParams{DataSourceType: "prometheus", Debug: true})
Outer:
	for _, pt := range tg.PromqlExprTests {
		result, _, err := queries.Query(context.TODO(), pt.Expr, time.Unix(int64(time.Duration(pt.EvalTime).Seconds()), 0))
		if err != nil {
			// todo query error
			checkErrs = append(checkErrs, fmt.Errorf("    expr: %q, time: %s, err: %w", pt.Expr,
				pt.EvalTime.String(), err))
			continue
		}
		var gotSamples []parsedSample

		if *result.SeriesFetched == 0 {
			if len(pt.ExpSamples) != 0 {
				// todo failed error
				checkErrs = append(checkErrs, fmt.Errorf("    expr: %q, time: %s, err: %w", pt.Expr,
					pt.EvalTime.String(), err))
			}
			continue
		}
		for _, s := range result.Data {
			gotSamples = append(gotSamples, parsedSample{
				Labels: s.Labels,
				Value:  s.Values[0],
			})
		}
		var expSamples []parsedSample
		for _, s := range pt.ExpSamples {
			lb, err := promparser.ParseMetric(s.Labels)
			if err != nil {
				err = fmt.Errorf("labels %q: %w", s.Labels, err)
				checkErrs = append(checkErrs, fmt.Errorf("    expr: %q, time: %s, err: %w", pt.Expr,
					pt.EvalTime.String(), err))
				continue Outer
			}
			var expLb labels
			for _, l := range lb {
				expLb = append(expLb, datasource.Label{
					Name:  l.Name,
					Value: l.Value,
				})
			}
			expSamples = append(expSamples, parsedSample{
				Labels: expLb,
				Value:  s.Value,
			})

			sort.Slice(expSamples, func(i, j int) bool {
				return labelCompare(expSamples[i].Labels, expSamples[j].Labels) <= 0
			})
			sort.Slice(gotSamples, func(i, j int) bool {
				return labelCompare(gotSamples[i].Labels, gotSamples[j].Labels) <= 0
			})
			if !reflect.DeepEqual(expSamples, gotSamples) {
				checkErrs = append(checkErrs, fmt.Errorf("    expr: %q, time: %s,\n        exp: %v\n        got: %v", pt.Expr,
					pt.EvalTime.String(), parsedSamplesString(expSamples), parsedSamplesString(gotSamples)))
			}
		}

	}

	logger.Infof("will sleep here %s", "wang")

	// wait here so we can using vmui
	for {
		select {
		case s := <-sigs:
			logger.Infof("program will exit now cause receiving signal %s", s)
			os.Exit(1)
		default:
		}
	}

	// todo check result
	// todo check promql_expr
	return checkErrs
}

func parsedSamplesString(pss []parsedSample) string {
	if len(pss) == 0 {
		return "nil"
	}
	s := pss[0].String()
	for _, ps := range pss[1:] {
		s += ", " + ps.String()
	}
	return s
}

// parsedSample is a sample with parsed Labels.
type parsedSample struct {
	Labels labels
	Value  float64
}

func (ps *parsedSample) String() string {
	return ps.Labels.String() + " " + strconv.FormatFloat(ps.Value, 'E', -1, 64)
}

// indentLines prefixes each line in the supplied string with the given "indent"
// string.
func indentLines(lines, indent string) string {
	sb := strings.Builder{}
	n := strings.Split(lines, "\n")
	for i, l := range n {
		if i > 0 {
			sb.WriteString(indent)
		}
		sb.WriteString(l)
		if i != len(n)-1 {
			sb.WriteRune('\n')
		}
	}
	return sb.String()
}

func convertToLabels(m map[string]string) (labelset labels) {
	for k, v := range m {
		labelset = append(labelset, datasource.Label{
			Name:  k,
			Value: v,
		})
	}
	return
}

type labels []datasource.Label

func (ls labels) Len() int           { return len(ls) }
func (ls labels) Swap(i, j int)      { ls[i], ls[j] = ls[j], ls[i] }
func (ls labels) Less(i, j int) bool { return ls[i].Name < ls[j].Name }

func (ls labels) String() string {
	var b bytes.Buffer

	b.WriteByte('{')
	for i, l := range ls {
		if i > 0 {
			b.WriteByte(',')
			b.WriteByte(' ')
		}
		b.WriteString(l.Name)
		b.WriteByte('=')
		b.WriteString(strconv.Quote(l.Value))
	}
	b.WriteByte('}')
	return b.String()
}

type labelAndAnnotation struct {
	Labels      labels
	Annotations labels
}

func (la *labelAndAnnotation) String() string {
	return "Labels:" + la.Labels.String() + "\nAnnotations:" + la.Annotations.String()
}

type labelsAndAnnotations []labelAndAnnotation

func (la labelsAndAnnotations) Len() int { return len(la) }

func (la labelsAndAnnotations) Swap(i, j int) { la[i], la[j] = la[j], la[i] }
func (la labelsAndAnnotations) Less(i, j int) bool {
	diff := labelCompare(la[i].Labels, la[j].Labels)
	if diff != 0 {
		return diff < 0
	}
	return labelCompare(la[i].Annotations, la[j].Annotations) < 0
}

func labelCompare(a, b labels) int {
	l := len(a)
	if len(b) < l {
		l = len(b)
	}

	for i := 0; i < l; i++ {
		if a[i].Name != b[i].Name {
			if a[i].Name < b[i].Name {
				return -1
			}
			return 1
		}
		if a[i].Value != b[i].Value {
			if a[i].Value < b[i].Value {
				return -1
			}
			return 1
		}
	}
	// If all labels so far were in common, the set with fewer labels comes first.
	return len(a) - len(b)
}

func (la labelsAndAnnotations) String() string {
	if len(la) == 0 {
		return "[]"
	}
	s := "[\n0:" + indentLines("\n"+la[0].String(), "  ")
	for i, l := range la[1:] {
		s += ",\n" + fmt.Sprintf("%d", i+1) + ":" + indentLines("\n"+l.String(), "  ")
	}
	s += "\n]"

	return s
}
