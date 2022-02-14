// Code generated by qtc from "web.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line app/vmalert/web.qtpl:1
package main

//line app/vmalert/web.qtpl:3
import (
	"path"
	"sort"
	"time"

	"github.com/VictoriaMetrics/VictoriaMetrics/app/vmalert/notifier"
	"github.com/VictoriaMetrics/VictoriaMetrics/app/vmalert/tpl"
)

//line app/vmalert/web.qtpl:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line app/vmalert/web.qtpl:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line app/vmalert/web.qtpl:13
func StreamWelcome(qw422016 *qt422016.Writer) {
//line app/vmalert/web.qtpl:13
	qw422016.N().S(`
    `)
//line app/vmalert/web.qtpl:14
	tpl.StreamHeader(qw422016, "vmalert", navItems)
//line app/vmalert/web.qtpl:14
	qw422016.N().S(`
    <p>
        API:<br>
        `)
//line app/vmalert/web.qtpl:17
	for _, p := range apiLinks {
//line app/vmalert/web.qtpl:17
		qw422016.N().S(`
            `)
//line app/vmalert/web.qtpl:19
		p, doc := p[0], p[1]

//line app/vmalert/web.qtpl:20
		qw422016.N().S(`
        	<a href="`)
//line app/vmalert/web.qtpl:21
		qw422016.E().S(p)
//line app/vmalert/web.qtpl:21
		qw422016.N().S(`">`)
//line app/vmalert/web.qtpl:21
		qw422016.E().S(p)
//line app/vmalert/web.qtpl:21
		qw422016.N().S(`</a> - `)
//line app/vmalert/web.qtpl:21
		qw422016.E().S(doc)
//line app/vmalert/web.qtpl:21
		qw422016.N().S(`<br/>
        `)
//line app/vmalert/web.qtpl:22
	}
//line app/vmalert/web.qtpl:22
	qw422016.N().S(`
    </p>
    `)
//line app/vmalert/web.qtpl:24
	tpl.StreamFooter(qw422016)
//line app/vmalert/web.qtpl:24
	qw422016.N().S(`
`)
//line app/vmalert/web.qtpl:25
}

//line app/vmalert/web.qtpl:25
func WriteWelcome(qq422016 qtio422016.Writer) {
//line app/vmalert/web.qtpl:25
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmalert/web.qtpl:25
	StreamWelcome(qw422016)
//line app/vmalert/web.qtpl:25
	qt422016.ReleaseWriter(qw422016)
//line app/vmalert/web.qtpl:25
}

//line app/vmalert/web.qtpl:25
func Welcome() string {
//line app/vmalert/web.qtpl:25
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmalert/web.qtpl:25
	WriteWelcome(qb422016)
//line app/vmalert/web.qtpl:25
	qs422016 := string(qb422016.B)
//line app/vmalert/web.qtpl:25
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmalert/web.qtpl:25
	return qs422016
//line app/vmalert/web.qtpl:25
}

//line app/vmalert/web.qtpl:27
func StreamListGroups(qw422016 *qt422016.Writer, groups []APIGroup) {
//line app/vmalert/web.qtpl:27
	qw422016.N().S(`
    `)
//line app/vmalert/web.qtpl:28
	tpl.StreamHeader(qw422016, "Groups", navItems)
//line app/vmalert/web.qtpl:28
	qw422016.N().S(`
    `)
//line app/vmalert/web.qtpl:29
	if len(groups) > 0 {
//line app/vmalert/web.qtpl:29
		qw422016.N().S(`
        `)
//line app/vmalert/web.qtpl:31
		rOk := make(map[string]int)
		rNotOk := make(map[string]int)
		for _, g := range groups {
			for _, r := range g.AlertingRules {
				if r.LastError != "" {
					rNotOk[g.Name]++
				} else {
					rOk[g.Name]++
				}
			}
			for _, r := range g.RecordingRules {
				if r.LastError != "" {
					rNotOk[g.Name]++
				} else {
					rOk[g.Name]++
				}
			}
		}

//line app/vmalert/web.qtpl:49
		qw422016.N().S(`
         <a class="btn btn-primary" role="button" onclick="collapseAll()">Collapse All</a>
         <a class="btn btn-primary" role="button" onclick="expandAll()">Expand All</a>
        `)
//line app/vmalert/web.qtpl:52
		for _, g := range groups {
//line app/vmalert/web.qtpl:52
			qw422016.N().S(`
              <div class="group-heading`)
//line app/vmalert/web.qtpl:53
			if rNotOk[g.Name] > 0 {
//line app/vmalert/web.qtpl:53
				qw422016.N().S(` alert-danger`)
//line app/vmalert/web.qtpl:53
			}
//line app/vmalert/web.qtpl:53
			qw422016.N().S(`"  data-bs-target="rules-`)
//line app/vmalert/web.qtpl:53
			qw422016.E().S(g.ID)
//line app/vmalert/web.qtpl:53
			qw422016.N().S(`">
                <span class="anchor" id="group-`)
//line app/vmalert/web.qtpl:54
			qw422016.E().S(g.ID)
//line app/vmalert/web.qtpl:54
			qw422016.N().S(`"></span>
                <a href="#group-`)
//line app/vmalert/web.qtpl:55
			qw422016.E().S(g.ID)
//line app/vmalert/web.qtpl:55
			qw422016.N().S(`">`)
//line app/vmalert/web.qtpl:55
			qw422016.E().S(g.Name)
//line app/vmalert/web.qtpl:55
			if g.Type != "prometheus" {
//line app/vmalert/web.qtpl:55
				qw422016.N().S(` (`)
//line app/vmalert/web.qtpl:55
				qw422016.E().S(g.Type)
//line app/vmalert/web.qtpl:55
				qw422016.N().S(`)`)
//line app/vmalert/web.qtpl:55
			}
//line app/vmalert/web.qtpl:55
			qw422016.N().S(` (every `)
//line app/vmalert/web.qtpl:55
			qw422016.E().S(g.Interval)
//line app/vmalert/web.qtpl:55
			qw422016.N().S(`)</a>
                 `)
//line app/vmalert/web.qtpl:56
			if rNotOk[g.Name] > 0 {
//line app/vmalert/web.qtpl:56
				qw422016.N().S(`<span class="badge bg-danger" title="Number of rules with status Error">`)
//line app/vmalert/web.qtpl:56
				qw422016.N().D(rNotOk[g.Name])
//line app/vmalert/web.qtpl:56
				qw422016.N().S(`</span> `)
//line app/vmalert/web.qtpl:56
			}
//line app/vmalert/web.qtpl:56
			qw422016.N().S(`
                <span class="badge bg-success" title="Number of rules withs status Ok">`)
//line app/vmalert/web.qtpl:57
			qw422016.N().D(rOk[g.Name])
//line app/vmalert/web.qtpl:57
			qw422016.N().S(`</span>
                <p class="fs-6 fw-lighter">`)
//line app/vmalert/web.qtpl:58
			qw422016.E().S(g.File)
//line app/vmalert/web.qtpl:58
			qw422016.N().S(`</p>
                `)
//line app/vmalert/web.qtpl:59
			if len(g.Params) > 0 {
//line app/vmalert/web.qtpl:59
				qw422016.N().S(`
                    <div class="fs-6 fw-lighter">Extra params
                    `)
//line app/vmalert/web.qtpl:61
				for _, param := range g.Params {
//line app/vmalert/web.qtpl:61
					qw422016.N().S(`
                            <span class="float-left badge bg-primary">`)
//line app/vmalert/web.qtpl:62
					qw422016.E().S(param)
//line app/vmalert/web.qtpl:62
					qw422016.N().S(`</span>
                    `)
//line app/vmalert/web.qtpl:63
				}
//line app/vmalert/web.qtpl:63
				qw422016.N().S(`
                    </div>
                `)
//line app/vmalert/web.qtpl:65
			}
//line app/vmalert/web.qtpl:65
			qw422016.N().S(`
            </div>
            <div class="collapse" id="rules-`)
//line app/vmalert/web.qtpl:67
			qw422016.E().S(g.ID)
//line app/vmalert/web.qtpl:67
			qw422016.N().S(`">
                <table class="table table-striped table-hover table-sm">
                    <thead>
                        <tr>
                            <th scope="col">Rule</th>
                            <th scope="col" title="Shows if rule's execution ended with error">Error</th>
                            <th scope="col" title="How many samples were produced by the rule">Samples</th>
                            <th scope="col" title="How many seconds ago rule was executed">Updated</th>
                        </tr>
                    </thead>
                    <tbody>
                    `)
//line app/vmalert/web.qtpl:78
			for _, ar := range g.AlertingRules {
//line app/vmalert/web.qtpl:78
				qw422016.N().S(`
                        <tr`)
//line app/vmalert/web.qtpl:79
				if ar.LastError != "" {
//line app/vmalert/web.qtpl:79
					qw422016.N().S(` class="alert-danger"`)
//line app/vmalert/web.qtpl:79
				}
//line app/vmalert/web.qtpl:79
				qw422016.N().S(`>
                            <td>
                                <b>alert:</b> `)
//line app/vmalert/web.qtpl:81
				qw422016.E().S(ar.Name)
//line app/vmalert/web.qtpl:81
				qw422016.N().S(` (for: `)
//line app/vmalert/web.qtpl:81
				qw422016.E().V(ar.For)
//line app/vmalert/web.qtpl:81
				qw422016.N().S(`)<br>
                                <code><pre>`)
//line app/vmalert/web.qtpl:82
				qw422016.E().S(ar.Expression)
//line app/vmalert/web.qtpl:82
				qw422016.N().S(`</pre></code><br>
                                `)
//line app/vmalert/web.qtpl:83
				if len(ar.Labels) > 0 {
//line app/vmalert/web.qtpl:83
					qw422016.N().S(` <b>Labels:</b>`)
//line app/vmalert/web.qtpl:83
				}
//line app/vmalert/web.qtpl:83
				qw422016.N().S(`
                                `)
//line app/vmalert/web.qtpl:84
				for k, v := range ar.Labels {
//line app/vmalert/web.qtpl:84
					qw422016.N().S(`
                                        <span class="ms-1 badge bg-primary">`)
//line app/vmalert/web.qtpl:85
					qw422016.E().S(k)
//line app/vmalert/web.qtpl:85
					qw422016.N().S(`=`)
//line app/vmalert/web.qtpl:85
					qw422016.E().S(v)
//line app/vmalert/web.qtpl:85
					qw422016.N().S(`</span>
                                `)
//line app/vmalert/web.qtpl:86
				}
//line app/vmalert/web.qtpl:86
				qw422016.N().S(`
                            </td>
                            <td><div class="error-cell">`)
//line app/vmalert/web.qtpl:88
				qw422016.E().S(ar.LastError)
//line app/vmalert/web.qtpl:88
				qw422016.N().S(`</div></td>
                            <td>`)
//line app/vmalert/web.qtpl:89
				qw422016.N().D(ar.LastSamples)
//line app/vmalert/web.qtpl:89
				qw422016.N().S(`</td>
                            <td>`)
//line app/vmalert/web.qtpl:90
				qw422016.N().FPrec(time.Since(ar.LastExec).Seconds(), 3)
//line app/vmalert/web.qtpl:90
				qw422016.N().S(`s ago</td>
                        </tr>
                    `)
//line app/vmalert/web.qtpl:92
			}
//line app/vmalert/web.qtpl:92
			qw422016.N().S(`
                    `)
//line app/vmalert/web.qtpl:93
			for _, rr := range g.RecordingRules {
//line app/vmalert/web.qtpl:93
				qw422016.N().S(`
                        <tr>
                            <td>
                                <b>record:</b> `)
//line app/vmalert/web.qtpl:96
				qw422016.E().S(rr.Name)
//line app/vmalert/web.qtpl:96
				qw422016.N().S(`<br>
                                <code><pre>`)
//line app/vmalert/web.qtpl:97
				qw422016.E().S(rr.Expression)
//line app/vmalert/web.qtpl:97
				qw422016.N().S(`</pre></code>
                                `)
//line app/vmalert/web.qtpl:98
				if len(rr.Labels) > 0 {
//line app/vmalert/web.qtpl:98
					qw422016.N().S(` <b>Labels:</b>`)
//line app/vmalert/web.qtpl:98
				}
//line app/vmalert/web.qtpl:98
				qw422016.N().S(`
                                `)
//line app/vmalert/web.qtpl:99
				for k, v := range rr.Labels {
//line app/vmalert/web.qtpl:99
					qw422016.N().S(`
                                        <span class="ms-1 badge bg-primary">`)
//line app/vmalert/web.qtpl:100
					qw422016.E().S(k)
//line app/vmalert/web.qtpl:100
					qw422016.N().S(`=`)
//line app/vmalert/web.qtpl:100
					qw422016.E().S(v)
//line app/vmalert/web.qtpl:100
					qw422016.N().S(`</span>
                                `)
//line app/vmalert/web.qtpl:101
				}
//line app/vmalert/web.qtpl:101
				qw422016.N().S(`
                            </td>
                            <td><div class="error-cell">`)
//line app/vmalert/web.qtpl:103
				qw422016.E().S(rr.LastError)
//line app/vmalert/web.qtpl:103
				qw422016.N().S(`</div></td>
                            <td>`)
//line app/vmalert/web.qtpl:104
				qw422016.N().D(rr.LastSamples)
//line app/vmalert/web.qtpl:104
				qw422016.N().S(`</td>
                            <td>`)
//line app/vmalert/web.qtpl:105
				qw422016.N().FPrec(time.Since(rr.LastExec).Seconds(), 3)
//line app/vmalert/web.qtpl:105
				qw422016.N().S(`s ago</td>
                        </tr>
                    `)
//line app/vmalert/web.qtpl:107
			}
//line app/vmalert/web.qtpl:107
			qw422016.N().S(`
                 </tbody>
                </table>
            </div>
        `)
//line app/vmalert/web.qtpl:111
		}
//line app/vmalert/web.qtpl:111
		qw422016.N().S(`

    `)
//line app/vmalert/web.qtpl:113
	} else {
//line app/vmalert/web.qtpl:113
		qw422016.N().S(`
        <div>
            <p>No items...</p>
        </div>
    `)
//line app/vmalert/web.qtpl:117
	}
//line app/vmalert/web.qtpl:117
	qw422016.N().S(`

    `)
//line app/vmalert/web.qtpl:119
	tpl.StreamFooter(qw422016)
//line app/vmalert/web.qtpl:119
	qw422016.N().S(`

`)
//line app/vmalert/web.qtpl:121
}

//line app/vmalert/web.qtpl:121
func WriteListGroups(qq422016 qtio422016.Writer, groups []APIGroup) {
//line app/vmalert/web.qtpl:121
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmalert/web.qtpl:121
	StreamListGroups(qw422016, groups)
//line app/vmalert/web.qtpl:121
	qt422016.ReleaseWriter(qw422016)
//line app/vmalert/web.qtpl:121
}

//line app/vmalert/web.qtpl:121
func ListGroups(groups []APIGroup) string {
//line app/vmalert/web.qtpl:121
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmalert/web.qtpl:121
	WriteListGroups(qb422016, groups)
//line app/vmalert/web.qtpl:121
	qs422016 := string(qb422016.B)
//line app/vmalert/web.qtpl:121
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmalert/web.qtpl:121
	return qs422016
//line app/vmalert/web.qtpl:121
}

//line app/vmalert/web.qtpl:124
func StreamListAlerts(qw422016 *qt422016.Writer, pathPrefix string, groupAlerts []GroupAlerts) {
//line app/vmalert/web.qtpl:124
	qw422016.N().S(`
    `)
//line app/vmalert/web.qtpl:125
	tpl.StreamHeader(qw422016, "Alerts", navItems)
//line app/vmalert/web.qtpl:125
	qw422016.N().S(`
    `)
//line app/vmalert/web.qtpl:126
	if len(groupAlerts) > 0 {
//line app/vmalert/web.qtpl:126
		qw422016.N().S(`
         <a class="btn btn-primary" role="button" onclick="collapseAll()">Collapse All</a>
         <a class="btn btn-primary" role="button" onclick="expandAll()">Expand All</a>
         `)
//line app/vmalert/web.qtpl:129
		for _, ga := range groupAlerts {
//line app/vmalert/web.qtpl:129
			qw422016.N().S(`
            `)
//line app/vmalert/web.qtpl:130
			g := ga.Group

//line app/vmalert/web.qtpl:130
			qw422016.N().S(`
            <div class="group-heading alert-danger" data-bs-target="rules-`)
//line app/vmalert/web.qtpl:131
			qw422016.E().S(g.ID)
//line app/vmalert/web.qtpl:131
			qw422016.N().S(`">
                <span class="anchor" id="group-`)
//line app/vmalert/web.qtpl:132
			qw422016.E().S(g.ID)
//line app/vmalert/web.qtpl:132
			qw422016.N().S(`"></span>
                <a href="#group-`)
//line app/vmalert/web.qtpl:133
			qw422016.E().S(g.ID)
//line app/vmalert/web.qtpl:133
			qw422016.N().S(`">`)
//line app/vmalert/web.qtpl:133
			qw422016.E().S(g.Name)
//line app/vmalert/web.qtpl:133
			if g.Type != "prometheus" {
//line app/vmalert/web.qtpl:133
				qw422016.N().S(` (`)
//line app/vmalert/web.qtpl:133
				qw422016.E().S(g.Type)
//line app/vmalert/web.qtpl:133
				qw422016.N().S(`)`)
//line app/vmalert/web.qtpl:133
			}
//line app/vmalert/web.qtpl:133
			qw422016.N().S(`</a>
                <span class="badge bg-danger" title="Number of active alerts">`)
//line app/vmalert/web.qtpl:134
			qw422016.N().D(len(ga.Alerts))
//line app/vmalert/web.qtpl:134
			qw422016.N().S(`</span>
                <br>
                <p class="fs-6 fw-lighter">`)
//line app/vmalert/web.qtpl:136
			qw422016.E().S(g.File)
//line app/vmalert/web.qtpl:136
			qw422016.N().S(`</p>
            </div>
            `)
//line app/vmalert/web.qtpl:139
			var keys []string
			alertsByRule := make(map[string][]*APIAlert)
			for _, alert := range ga.Alerts {
				if len(alertsByRule[alert.RuleID]) < 1 {
					keys = append(keys, alert.RuleID)
				}
				alertsByRule[alert.RuleID] = append(alertsByRule[alert.RuleID], alert)
			}
			sort.Strings(keys)

//line app/vmalert/web.qtpl:148
			qw422016.N().S(`
            <div class="collapse" id="rules-`)
//line app/vmalert/web.qtpl:149
			qw422016.E().S(g.ID)
//line app/vmalert/web.qtpl:149
			qw422016.N().S(`">
                `)
//line app/vmalert/web.qtpl:150
			for _, ruleID := range keys {
//line app/vmalert/web.qtpl:150
				qw422016.N().S(`
                    `)
//line app/vmalert/web.qtpl:152
				defaultAR := alertsByRule[ruleID][0]
				var labelKeys []string
				for k := range defaultAR.Labels {
					labelKeys = append(labelKeys, k)
				}
				sort.Strings(labelKeys)

//line app/vmalert/web.qtpl:158
				qw422016.N().S(`
                    <br>
                    <b>alert:</b> `)
//line app/vmalert/web.qtpl:160
				qw422016.E().S(defaultAR.Name)
//line app/vmalert/web.qtpl:160
				qw422016.N().S(` (`)
//line app/vmalert/web.qtpl:160
				qw422016.N().D(len(alertsByRule[ruleID]))
//line app/vmalert/web.qtpl:160
				qw422016.N().S(`)
                     | <span><a target="_blank" href="`)
//line app/vmalert/web.qtpl:161
				qw422016.E().S(defaultAR.SourceLink)
//line app/vmalert/web.qtpl:161
				qw422016.N().S(`">Source</a></span>
                    <br>
                    <b>expr:</b><code><pre>`)
//line app/vmalert/web.qtpl:163
				qw422016.E().S(defaultAR.Expression)
//line app/vmalert/web.qtpl:163
				qw422016.N().S(`</pre></code>
                    <table class="table table-striped table-hover table-sm">
                        <thead>
                            <tr>
                                <th scope="col">Labels</th>
                                <th scope="col">State</th>
                                <th scope="col">Active at</th>
                                <th scope="col">Value</th>
                                <th scope="col">Link</th>
                            </tr>
                        </thead>
                        <tbody>
                        `)
//line app/vmalert/web.qtpl:175
				for _, ar := range alertsByRule[ruleID] {
//line app/vmalert/web.qtpl:175
					qw422016.N().S(`
                            <tr>
                                <td>
                                    `)
//line app/vmalert/web.qtpl:178
					for _, k := range labelKeys {
//line app/vmalert/web.qtpl:178
						qw422016.N().S(`
                                        <span class="ms-1 badge bg-primary">`)
//line app/vmalert/web.qtpl:179
						qw422016.E().S(k)
//line app/vmalert/web.qtpl:179
						qw422016.N().S(`=`)
//line app/vmalert/web.qtpl:179
						qw422016.E().S(ar.Labels[k])
//line app/vmalert/web.qtpl:179
						qw422016.N().S(`</span>
                                    `)
//line app/vmalert/web.qtpl:180
					}
//line app/vmalert/web.qtpl:180
					qw422016.N().S(`
                                </td>
                                <td>`)
//line app/vmalert/web.qtpl:182
					streambadgeState(qw422016, ar.State)
//line app/vmalert/web.qtpl:182
					qw422016.N().S(`</td>
                                <td>
                                    `)
//line app/vmalert/web.qtpl:184
					qw422016.E().S(ar.ActiveAt.Format("2006-01-02T15:04:05Z07:00"))
//line app/vmalert/web.qtpl:184
					qw422016.N().S(`
                                    `)
//line app/vmalert/web.qtpl:185
					if ar.Restored {
//line app/vmalert/web.qtpl:185
						streambadgeRestored(qw422016)
//line app/vmalert/web.qtpl:185
					}
//line app/vmalert/web.qtpl:185
					qw422016.N().S(`
                                </td>
                                <td>`)
//line app/vmalert/web.qtpl:187
					qw422016.E().S(ar.Value)
//line app/vmalert/web.qtpl:187
					qw422016.N().S(`</td>
                                <td>
                                    <a href="`)
//line app/vmalert/web.qtpl:189
					qw422016.E().S(path.Join(pathPrefix, g.ID, ar.ID, "status"))
//line app/vmalert/web.qtpl:189
					qw422016.N().S(`">Details</a>
                                </td>
                            </tr>
                        `)
//line app/vmalert/web.qtpl:192
				}
//line app/vmalert/web.qtpl:192
				qw422016.N().S(`
                     </tbody>
                    </table>
                `)
//line app/vmalert/web.qtpl:195
			}
//line app/vmalert/web.qtpl:195
			qw422016.N().S(`
            </div>
            <br>
        `)
//line app/vmalert/web.qtpl:198
		}
//line app/vmalert/web.qtpl:198
		qw422016.N().S(`

    `)
//line app/vmalert/web.qtpl:200
	} else {
//line app/vmalert/web.qtpl:200
		qw422016.N().S(`
        <div>
            <p>No items...</p>
        </div>
    `)
//line app/vmalert/web.qtpl:204
	}
//line app/vmalert/web.qtpl:204
	qw422016.N().S(`

    `)
//line app/vmalert/web.qtpl:206
	tpl.StreamFooter(qw422016)
//line app/vmalert/web.qtpl:206
	qw422016.N().S(`

`)
//line app/vmalert/web.qtpl:208
}

//line app/vmalert/web.qtpl:208
func WriteListAlerts(qq422016 qtio422016.Writer, pathPrefix string, groupAlerts []GroupAlerts) {
//line app/vmalert/web.qtpl:208
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmalert/web.qtpl:208
	StreamListAlerts(qw422016, pathPrefix, groupAlerts)
//line app/vmalert/web.qtpl:208
	qt422016.ReleaseWriter(qw422016)
//line app/vmalert/web.qtpl:208
}

//line app/vmalert/web.qtpl:208
func ListAlerts(pathPrefix string, groupAlerts []GroupAlerts) string {
//line app/vmalert/web.qtpl:208
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmalert/web.qtpl:208
	WriteListAlerts(qb422016, pathPrefix, groupAlerts)
//line app/vmalert/web.qtpl:208
	qs422016 := string(qb422016.B)
//line app/vmalert/web.qtpl:208
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmalert/web.qtpl:208
	return qs422016
//line app/vmalert/web.qtpl:208
}

//line app/vmalert/web.qtpl:210
func StreamListTargets(qw422016 *qt422016.Writer, targets map[notifier.TargetType][]notifier.Target) {
//line app/vmalert/web.qtpl:210
	qw422016.N().S(`
    `)
//line app/vmalert/web.qtpl:211
	tpl.StreamHeader(qw422016, "Notifiers", navItems)
//line app/vmalert/web.qtpl:211
	qw422016.N().S(`
    `)
//line app/vmalert/web.qtpl:212
	if len(targets) > 0 {
//line app/vmalert/web.qtpl:212
		qw422016.N().S(`
         <a class="btn btn-primary" role="button" onclick="collapseAll()">Collapse All</a>
         <a class="btn btn-primary" role="button" onclick="expandAll()">Expand All</a>

         `)
//line app/vmalert/web.qtpl:217
		var keys []string
		for key := range targets {
			keys = append(keys, string(key))
		}
		sort.Strings(keys)

//line app/vmalert/web.qtpl:222
		qw422016.N().S(`

         `)
//line app/vmalert/web.qtpl:224
		for i := range keys {
//line app/vmalert/web.qtpl:224
			qw422016.N().S(`
           `)
//line app/vmalert/web.qtpl:225
			typeK, ns := keys[i], targets[notifier.TargetType(keys[i])]
			count := len(ns)

//line app/vmalert/web.qtpl:227
			qw422016.N().S(`
           <div class="group-heading data-bs-target="rules-`)
//line app/vmalert/web.qtpl:228
			qw422016.E().S(typeK)
//line app/vmalert/web.qtpl:228
			qw422016.N().S(`">
             <span class="anchor" id="notifiers-`)
//line app/vmalert/web.qtpl:229
			qw422016.E().S(typeK)
//line app/vmalert/web.qtpl:229
			qw422016.N().S(`"></span>
             <a href="#notifiers-`)
//line app/vmalert/web.qtpl:230
			qw422016.E().S(typeK)
//line app/vmalert/web.qtpl:230
			qw422016.N().S(`">`)
//line app/vmalert/web.qtpl:230
			qw422016.E().S(typeK)
//line app/vmalert/web.qtpl:230
			qw422016.N().S(` (`)
//line app/vmalert/web.qtpl:230
			qw422016.N().D(count)
//line app/vmalert/web.qtpl:230
			qw422016.N().S(`)</a>
         </div>
         <div class="collapse show" id="notifiers-`)
//line app/vmalert/web.qtpl:232
			qw422016.E().S(typeK)
//line app/vmalert/web.qtpl:232
			qw422016.N().S(`">
             <table class="table table-striped table-hover table-sm">
                 <thead>
                     <tr>
                         <th scope="col">Labels</th>
                         <th scope="col">Address</th>
                     </tr>
                 </thead>
                 <tbody>
                 `)
//line app/vmalert/web.qtpl:241
			for _, n := range ns {
//line app/vmalert/web.qtpl:241
				qw422016.N().S(`
                     <tr>
                         <td>
                              `)
//line app/vmalert/web.qtpl:244
				for _, l := range n.Labels {
//line app/vmalert/web.qtpl:244
					qw422016.N().S(`
                                      <span class="ms-1 badge bg-primary">`)
//line app/vmalert/web.qtpl:245
					qw422016.E().S(l.Name)
//line app/vmalert/web.qtpl:245
					qw422016.N().S(`=`)
//line app/vmalert/web.qtpl:245
					qw422016.E().S(l.Value)
//line app/vmalert/web.qtpl:245
					qw422016.N().S(`</span>
                              `)
//line app/vmalert/web.qtpl:246
				}
//line app/vmalert/web.qtpl:246
				qw422016.N().S(`
                          </td>
                         <td>`)
//line app/vmalert/web.qtpl:248
				qw422016.E().S(n.Notifier.Addr())
//line app/vmalert/web.qtpl:248
				qw422016.N().S(`</td>
                     </tr>
                 `)
//line app/vmalert/web.qtpl:250
			}
//line app/vmalert/web.qtpl:250
			qw422016.N().S(`
              </tbody>
             </table>
         </div>
     `)
//line app/vmalert/web.qtpl:254
		}
//line app/vmalert/web.qtpl:254
		qw422016.N().S(`

    `)
//line app/vmalert/web.qtpl:256
	} else {
//line app/vmalert/web.qtpl:256
		qw422016.N().S(`
        <div>
            <p>No items...</p>
        </div>
    `)
//line app/vmalert/web.qtpl:260
	}
//line app/vmalert/web.qtpl:260
	qw422016.N().S(`

    `)
//line app/vmalert/web.qtpl:262
	tpl.StreamFooter(qw422016)
//line app/vmalert/web.qtpl:262
	qw422016.N().S(`

`)
//line app/vmalert/web.qtpl:264
}

//line app/vmalert/web.qtpl:264
func WriteListTargets(qq422016 qtio422016.Writer, targets map[notifier.TargetType][]notifier.Target) {
//line app/vmalert/web.qtpl:264
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmalert/web.qtpl:264
	StreamListTargets(qw422016, targets)
//line app/vmalert/web.qtpl:264
	qt422016.ReleaseWriter(qw422016)
//line app/vmalert/web.qtpl:264
}

//line app/vmalert/web.qtpl:264
func ListTargets(targets map[notifier.TargetType][]notifier.Target) string {
//line app/vmalert/web.qtpl:264
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmalert/web.qtpl:264
	WriteListTargets(qb422016, targets)
//line app/vmalert/web.qtpl:264
	qs422016 := string(qb422016.B)
//line app/vmalert/web.qtpl:264
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmalert/web.qtpl:264
	return qs422016
//line app/vmalert/web.qtpl:264
}

//line app/vmalert/web.qtpl:266
func StreamAlert(qw422016 *qt422016.Writer, pathPrefix string, alert *APIAlert) {
//line app/vmalert/web.qtpl:266
	qw422016.N().S(`
    `)
//line app/vmalert/web.qtpl:267
	tpl.StreamHeader(qw422016, "", navItems)
//line app/vmalert/web.qtpl:267
	qw422016.N().S(`
    `)
//line app/vmalert/web.qtpl:269
	var labelKeys []string
	for k := range alert.Labels {
		labelKeys = append(labelKeys, k)
	}
	sort.Strings(labelKeys)

	var annotationKeys []string
	for k := range alert.Annotations {
		annotationKeys = append(annotationKeys, k)
	}
	sort.Strings(annotationKeys)

//line app/vmalert/web.qtpl:280
	qw422016.N().S(`
    <div class="display-6 pb-3 mb-3">`)
//line app/vmalert/web.qtpl:281
	qw422016.E().S(alert.Name)
//line app/vmalert/web.qtpl:281
	qw422016.N().S(`<span class="ms-2 badge `)
//line app/vmalert/web.qtpl:281
	if alert.State == "firing" {
//line app/vmalert/web.qtpl:281
		qw422016.N().S(`bg-danger`)
//line app/vmalert/web.qtpl:281
	} else {
//line app/vmalert/web.qtpl:281
		qw422016.N().S(` bg-warning text-dark`)
//line app/vmalert/web.qtpl:281
	}
//line app/vmalert/web.qtpl:281
	qw422016.N().S(`">`)
//line app/vmalert/web.qtpl:281
	qw422016.E().S(alert.State)
//line app/vmalert/web.qtpl:281
	qw422016.N().S(`</span></div>
    <div class="container border-bottom p-2">
      <div class="row">
        <div class="col-2">
          Active at
        </div>
        <div class="col">
          `)
//line app/vmalert/web.qtpl:288
	qw422016.E().S(alert.ActiveAt.Format("2006-01-02T15:04:05Z07:00"))
//line app/vmalert/web.qtpl:288
	qw422016.N().S(`
        </div>
      </div>
      </div>
    <div class="container border-bottom p-2">
      <div class="row">
        <div class="col-2">
          Expr
        </div>
        <div class="col">
          <code><pre>`)
//line app/vmalert/web.qtpl:298
	qw422016.E().S(alert.Expression)
//line app/vmalert/web.qtpl:298
	qw422016.N().S(`</pre></code>
        </div>
      </div>
    </div>
    <div class="container border-bottom p-2">
      <div class="row">
        <div class="col-2">
          Labels
        </div>
        <div class="col">
           `)
//line app/vmalert/web.qtpl:308
	for _, k := range labelKeys {
//line app/vmalert/web.qtpl:308
		qw422016.N().S(`
                <span class="m-1 badge bg-primary">`)
//line app/vmalert/web.qtpl:309
		qw422016.E().S(k)
//line app/vmalert/web.qtpl:309
		qw422016.N().S(`=`)
//line app/vmalert/web.qtpl:309
		qw422016.E().S(alert.Labels[k])
//line app/vmalert/web.qtpl:309
		qw422016.N().S(`</span>
          `)
//line app/vmalert/web.qtpl:310
	}
//line app/vmalert/web.qtpl:310
	qw422016.N().S(`
        </div>
      </div>
    </div>
    <div class="container border-bottom p-2">
      <div class="row">
        <div class="col-2">
          Annotations
        </div>
        <div class="col">
           `)
//line app/vmalert/web.qtpl:320
	for _, k := range annotationKeys {
//line app/vmalert/web.qtpl:320
		qw422016.N().S(`
                <b>`)
//line app/vmalert/web.qtpl:321
		qw422016.E().S(k)
//line app/vmalert/web.qtpl:321
		qw422016.N().S(`:</b><br>
                <p>`)
//line app/vmalert/web.qtpl:322
		qw422016.E().S(alert.Annotations[k])
//line app/vmalert/web.qtpl:322
		qw422016.N().S(`</p>
          `)
//line app/vmalert/web.qtpl:323
	}
//line app/vmalert/web.qtpl:323
	qw422016.N().S(`
        </div>
      </div>
    </div>
    <div class="container border-bottom p-2">
      <div class="row">
        <div class="col-2">
          Group
        </div>
        <div class="col">
           <a target="_blank" href="`)
//line app/vmalert/web.qtpl:333
	qw422016.E().S(path.Join(pathPrefix, "groups"))
//line app/vmalert/web.qtpl:333
	qw422016.N().S(`#group-`)
//line app/vmalert/web.qtpl:333
	qw422016.E().S(alert.GroupID)
//line app/vmalert/web.qtpl:333
	qw422016.N().S(`">`)
//line app/vmalert/web.qtpl:333
	qw422016.E().S(alert.GroupID)
//line app/vmalert/web.qtpl:333
	qw422016.N().S(`</a>
        </div>
      </div>
    </div>
     <div class="container border-bottom p-2">
      <div class="row">
        <div class="col-2">
          Source link
        </div>
        <div class="col">
           <a target="_blank" href="`)
//line app/vmalert/web.qtpl:343
	qw422016.E().S(alert.SourceLink)
//line app/vmalert/web.qtpl:343
	qw422016.N().S(`">Link</a>
        </div>
      </div>
    </div>
    `)
//line app/vmalert/web.qtpl:347
	tpl.StreamFooter(qw422016)
//line app/vmalert/web.qtpl:347
	qw422016.N().S(`

`)
//line app/vmalert/web.qtpl:349
}

//line app/vmalert/web.qtpl:349
func WriteAlert(qq422016 qtio422016.Writer, pathPrefix string, alert *APIAlert) {
//line app/vmalert/web.qtpl:349
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmalert/web.qtpl:349
	StreamAlert(qw422016, pathPrefix, alert)
//line app/vmalert/web.qtpl:349
	qt422016.ReleaseWriter(qw422016)
//line app/vmalert/web.qtpl:349
}

//line app/vmalert/web.qtpl:349
func Alert(pathPrefix string, alert *APIAlert) string {
//line app/vmalert/web.qtpl:349
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmalert/web.qtpl:349
	WriteAlert(qb422016, pathPrefix, alert)
//line app/vmalert/web.qtpl:349
	qs422016 := string(qb422016.B)
//line app/vmalert/web.qtpl:349
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmalert/web.qtpl:349
	return qs422016
//line app/vmalert/web.qtpl:349
}

//line app/vmalert/web.qtpl:351
func streambadgeState(qw422016 *qt422016.Writer, state string) {
//line app/vmalert/web.qtpl:351
	qw422016.N().S(`
`)
//line app/vmalert/web.qtpl:353
	badgeClass := "bg-warning text-dark"
	if state == "firing" {
		badgeClass = "bg-danger"
	}

//line app/vmalert/web.qtpl:357
	qw422016.N().S(`
<span class="badge `)
//line app/vmalert/web.qtpl:358
	qw422016.E().S(badgeClass)
//line app/vmalert/web.qtpl:358
	qw422016.N().S(`">`)
//line app/vmalert/web.qtpl:358
	qw422016.E().S(state)
//line app/vmalert/web.qtpl:358
	qw422016.N().S(`</span>
`)
//line app/vmalert/web.qtpl:359
}

//line app/vmalert/web.qtpl:359
func writebadgeState(qq422016 qtio422016.Writer, state string) {
//line app/vmalert/web.qtpl:359
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmalert/web.qtpl:359
	streambadgeState(qw422016, state)
//line app/vmalert/web.qtpl:359
	qt422016.ReleaseWriter(qw422016)
//line app/vmalert/web.qtpl:359
}

//line app/vmalert/web.qtpl:359
func badgeState(state string) string {
//line app/vmalert/web.qtpl:359
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmalert/web.qtpl:359
	writebadgeState(qb422016, state)
//line app/vmalert/web.qtpl:359
	qs422016 := string(qb422016.B)
//line app/vmalert/web.qtpl:359
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmalert/web.qtpl:359
	return qs422016
//line app/vmalert/web.qtpl:359
}

//line app/vmalert/web.qtpl:361
func streambadgeRestored(qw422016 *qt422016.Writer) {
//line app/vmalert/web.qtpl:361
	qw422016.N().S(`
<span class="badge bg-warning text-dark" title="Alert state was restored after the service restart from remote storage">restored</span>
`)
//line app/vmalert/web.qtpl:363
}

//line app/vmalert/web.qtpl:363
func writebadgeRestored(qq422016 qtio422016.Writer) {
//line app/vmalert/web.qtpl:363
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmalert/web.qtpl:363
	streambadgeRestored(qw422016)
//line app/vmalert/web.qtpl:363
	qt422016.ReleaseWriter(qw422016)
//line app/vmalert/web.qtpl:363
}

//line app/vmalert/web.qtpl:363
func badgeRestored() string {
//line app/vmalert/web.qtpl:363
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmalert/web.qtpl:363
	writebadgeRestored(qb422016)
//line app/vmalert/web.qtpl:363
	qs422016 := string(qb422016.B)
//line app/vmalert/web.qtpl:363
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmalert/web.qtpl:363
	return qs422016
//line app/vmalert/web.qtpl:363
}
