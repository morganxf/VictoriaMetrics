## Metrics
### 采集失败Metrics
* vm_promscrape_scrapes_failed_total: 采集失败总量
    * vm_promscrape_scrapes_timed_out_total: 因超时采集失败总量
    * vm_promscrape_max_scrape_size_exceeded_errors_total: 因响应体过大采集失败总量，通过`-promscrape.maxScrapeSize`放宽限制
### 采集跳过Metrics
* vm_promscrape_scrapes_skipped_by_sample_limit_total: 因响应体序列数过多采样跳过总量，通过修改采集配置`sample_limit`参数解决
