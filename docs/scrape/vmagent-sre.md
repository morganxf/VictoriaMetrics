### 如何放宽对采集响应大小的限制
背景：vmagent默认只支持最大16M的响应体，如果Metrics响应体超过16M就会拒绝该采集

解决：vmagent CR增加`promscrape.maxScrapeSize` 参数，即`spec.extraArgs.promscrape.maxScrapeSize: 50MB`(设置最大限制50MB)
