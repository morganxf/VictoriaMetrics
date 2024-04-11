# 指标费用对比
简单结论：
* 阿里云Prometheus价格：保存90天，100百万数据点/天，价格552元/月
* 腾讯云Prometheus价格：保存90天计算，100百万数据点/天，价格4200元/月  
* Datadog价格：保存15月计算，100百万数据点/天，价格166666美元/月  


## Datadog计费
Datadog计费采用基本计费+单项计费联合计费模式。其中，基本计费是必须的，单项计费根据用户的实际使用情况收费。

Datadog计费特点：
1. 按关键功能划分不同的基本计费版本。
2. 按照产品功能单项计费，计费按照使用量计费（费用大头）。基本计费容量一定不够用
> * 按host、container、指标量等计费

3. 针对指标做了优化：间接索引计费、无索引低价、免费长时间数据存储

### 基本计费
基本计费提供了5种计费标准：
1. Free
2. Pro
3. Enterprise
4. DevSecOps Pro
5. DevSecOps Enterprise

| 模式 | 计费单位 | 计费 | 备注 |
| --- | --- | --- | --- |
| Free | | 0 | 指标保留1天，最多接入5个host(agent安装) |
| Pro | 每host每月 | 15$ | 免费5个容器、100个时间线(每host)；700+ integrations, 开箱即用大盘，指标保留15月 |
| Enterprise | 每host每月 | 23$ | 包含Pro功能；免费10个容器、200个时间线(每host)；基于机器学习的告警，进程监控 |
| DevSecOps Pro | 每host每月 | 22$ | 包含Pro功能；提供安全保护 |
| DevSecOps Enterprise | 每host每月 | 34$ | 包含Enterprise和DevSecOps Pro功能；提供风险和威胁检测功能 |

PS 上述内容做了裁剪，更详细内容见[Datadog价格](https://www.datadoghq.com/pricing/)。

### 单项计费

| 产品 | 计费单位 | 计费 |
| --- | --- | --- |
| 容器监控 | 每容器每月 | 1$ |
| 自定义指标(索引) | 每100时间线每月 | 5$ |
| 自定义指标(无索引) | 每100时间线每月 | 0.1$ |

* 针对自定义指标(索引)，每天100百万数据点，等价费用`166666$ =5.0/30*1000000` 
* 针对自定义指标(无索引)，每天100百万数据点，等价费用`3333$ =0.1/30*1000000` 


PS 上述内容做了裁剪，更详细内容见[Datadog价格列表](https://www.datadoghq.com/pricing/list/)

### 附录

* [Datadog价格](https://www.datadoghq.com/pricing/)
* [Datadog价格列表](https://www.datadoghq.com/pricing/list/)
* [Datadog计费](https://docs.datadoghq.com/account_management/billing/)
* [Datadog自定义指标计费](https://docs.datadoghq.com/account_management/billing/custom_metrics/?tab=countrate)
* [Datadog无限制指标](https://docs.datadoghq.com/metrics/metrics-without-limits/)

## 阿里云ARMS Prometheus计费
支持存储90天和180天两种，存储时长越长计费越贵。保存90天单价：0.4元/GB，保存180天单价：0.6元/GB。  
按照保存90天计算，每天100百万(10亿)个数据点，价格552元/月。ARMS根据线上数据统计，一条指标的平均大小为0.5Kb。

PS 阿里云云盘价格：
| 类型 | 价格 | IOPS |
| --- | --- | --- |
| SSD云盘 | 1元/GiB | 3300 |
| ESSD云盘 | 0.5元/GiB | 2400 |
| 高效云盘 | 0.35元/GiB | 2200 |


[ARMS价格计算器](https://armsnext.console.aliyun.com/price-gb?spm=5176.8140086.J_5253785160.7.3f29be45gWaq3d#/overview?product=prometheus)
[产品计费](https://help.aliyun.com/zh/arms/product-overview/product-billing-new-version?spm=5176.8140086.J_5253785160.6.3f29be45gWaq3d)

## 腾讯云Prometheus
支持多种存储时长，15天-2年不等，存储时长越长计费越贵。计算公式：  
日费用 = 日上报数据量（单位：百万条） × 单价  

按照保存90天计算，每天100百万(10亿)个数据点，价格4200元/月`=100 * 1.4 * 30` 

### 存储90天单价
| 日均上报数据量范围（百万条） | 单价（元） |
|----------------|-------|
| 0~200 | 1.4 |
| 200~800 | 1.2 |
| 800~1500 | 1.1 |
| 1500 以上 | 1 |

### 附录
* [按量计费](https://cloud.tencent.com/document/product/248/87091)
* [套餐包计费](https://cloud.tencent.com/document/product/248/87095)
* [Promethues优势](https://cloud.tencent.com/document/product/248/87372)


## 阿里云TSDB
阿里云提供两种TSDB，一是自研的Lindorm，一是InfluxDB。其中Lindorm不是专门的TSDB，只是能够提供等价TSDB的能力；InfluxDB是开源托管版，并提供高可用版(无法水平伸缩)。

### 云原生多模数据库Lindorm
Lindorm是面向物联网、互联网、车联网等设计和优化的云原生多模超融合数据库，支持宽表、时序、文本、对象、流、空间等多种数据的统一访问和融合处理，并兼容SQL、HBase/Cassandra/S3、TSDB、HDFS、Solr、Kafka等多种标准接口和无缝集成三方生态工具，适用于日志、监控、账单、广告、社交、出行、风控等场景。详见[Lindorm](https://help.aliyun.com/document_detail/174640.html?spm=a2c4g.174643.0.0.3bec23f6KJU66m)

Lindorm关键点：
1. 支持多数据模型
2. 开源兼容
3. 存储、计算分离
4. 各引擎、存储独立弹性容量
5. 冷、热存储

可以简单认为Lindorm是HBase+OpenTSDB+Elasticsearch+HDFS+Kafaka，详见[Lindorm引擎类型](https://help.aliyun.com/document_detail/174643.html?spm=a2c4g.174643.0.0.557123f6SOqiOd)

#### Lindrom TSDB计费 
| 实例规格 | 计费单位 | 计费 | 存储空间 | 备注 |
| --- | --- | --- | --- | --- |
| 基础版III | 月 | 3,732 | 320GB | 共享实例，最大接入数据测点数：1000万，TPS：4万 |
| 基础版II | 月 | 1,632 | 320GB | 共享实例，最大接入数据测点数：480万，TPS：3万 |
| 基础版I | 月 | 1,032 | 320GB | 共享实例，最大接入数据测点数：240万，TPS：1万 |
| 标准版II | 月 | 1,3376 | 320GB | 共享实例，最大接入数据测点数：4000万，TPS：24万 |
| 标准版I | 月 | 7,232 | 320GB | 共享实例，最大接入数据测点数：2000万，TPS：12万 |
| 旗舰版II | 月 | 48,090 | 320GB | 独享实例，最大接入数据测点数：无限制，TPS：96万 |
| 旗舰版I | 月 | 24,954 | 320GB | 独享实例，最大接入数据测点数：无限制，TPS：48万 |

#### 附录
* [Lindorm TSDB](https://www.aliyun.com/product/hitsdb?spm=5176.149792.J_XmGx2FZCDAeIy2ZCWL7sW.25.67f26ef0t9qeVL&scm=20140722.S_product@@%E4%BA%91%E4%BA%A7%E5%93%81@@72462._.RL_LindormTSDB-LOC_topbarproduct-OR_ser-V_3-RE_productNew-P0_0)
* [Lindorm](https://help.aliyun.com/document_detail/174640.html?spm=a2c4g.174643.0.0.3bec23f6KJU66m)
* [Lindorm引擎类型](https://help.aliyun.com/document_detail/174643.html?spm=a2c4g.174643.0.0.557123f6SOqiOd)
* [价格计算器](https://www.aliyun.com/price/product?spm=5176.149792.J_775960.price-tsdb.67f26ef0LBUgV7#/hitsdb/detail/hitsdbpre)

### InfluxDB 
托管InfluxDB，分为基础班和高可用版：
* 基础版：单点
* 高可用版：3节点集群，数据3副本

#### InfluxDB计费
InfluxDB价格基本等同ECS价格。以4C16G为例，高可用InfluxDB价格1503，3个ECS价格1506.96。

| 系列 | 基础版 | 基础版 | 基础版 | 基础版 | 基础版 | 高可用版 | 高可用版  | 高可用版 | 高可用版 | 高可用版 | 
|--------|--------|--------|---------|----------|----------|--------|---------|---------|----------|----------|
| 规格     | 4核16GB | 8核32GB | 16核64GB | 32核128GB | 64核256GB | 4核16GB | 8核32GB  | 16核64GB | 32核128GB | 64核256GB |
| 每秒写入请求 | ~100   | ~160   | ~300    | ~500     | ~800     | ~160   | ~240    | ~400    | ~600     | ~1000    |
| 写入数据点  | ~50000 | ~80000 | ~150000 | ~250000  | ~400000  | ~80000 | ~120000 | ~200000 | ~300000  | ~500000  |
| 每秒查询请求 | ~50    | ~80    | ~150    | ~250     | ~400     | ~170   | ~300    | ~480    | ~780     | ~1250    |
| 中国内地   | 541    | 1074    | 2140     | 4271     | 8535   | 1503    | 2986    | 5951     | 11881    | 23742 |

单位：元/GB/月

* 每秒写入请求指的是每秒写入请求的数量。
* 写入数据点为每秒写入的数据点数量。
* 查询性能为每秒可执行的查询次数，每次查询扫描10个时间线，1万个原始数据点。

## 腾讯云TSDB
CTSDB于2018年12月正式上线，一直在支撑腾讯云可观测平台。InfluxDB于2023年4月发布单机版，2024年2月上线集群版，目标是提供一款高性能读写海量时序数据库。

### CTSDB
CTSDB可以认为是针对TS优化后的Elastricsearch，数据存储是2副本。详见[CTSDB](https://cloud.tencent.com/document/product/652/13532)

#### 节点数量与规格
* 节点数量：配置至少3个节点。仅有2个节点的实例，无法完全保证数据的高可靠性，仅适用于测试环境。
* 节点规格：配置至少2核9GB。节点规格为1核2GB，内存过小无法保证高可用性，仅适用于测试环境。

#### 价格 
实例价格  
= 内存规格费用 + 存储空间费用  
= (总内存规格 * 单位内存价格) + (总存储空间 * 单位存储价格)  
= (单节点内存数 * 节点数 * 单位内存价格) + (单节点存储空间 * 节点数 * 单位存储价格)  
详见[产品定价](https://cloud.tencent.com/document/product/652/31942)

CTSDB 以2C9G350GB 3节点为例，实际价格1915.6，计算价格2187`=(67*9*3)+(0.36*350*3)`  
CVM 以2C8G350GB 3节点为例，实际价格价格1568.4

##### 内存价格
| CPU(核) | 内存(GB) | 适用场景 | 推荐写入次数 | 单位内存价格(元/GB/月) |
|-----|-----|--------|------|--------|
| 2 | 9 | 生产环境 | 50000 | 67 |
| 4 | 20 | 生产环境 | 100000 | 60.30 |
| 8 | 40 | 生产环境 | 200000 | 56.95 |
| 16 | 80 | 生产环境 | 400000 | 53.60 |
| 28 | 128 | 生产环境 | 800000 | 50.25 |

##### 存储价格
| 数据库类型 | 单位存储价格(元/GB/月)
|-----|-----|
| CTSDB实例 | 0.36 |

#### 附件
* [CTSDB](https://cloud.tencent.com/document/product/652/13532)
* [产品定价](https://cloud.tencent.com/document/product/652/31942)

### InfluxDB 
可弹性伸缩的InfluxDB集群，分时序节点和数据节点两种：[系统架构](https://cloud.tencent.com/document/product/652/89876)
1. 时序节点：计算节点。规格：可选，数量：[3, 30]
2. 数据节点：存储节点。规格：15C60GB，数量：9、18、27、36、45、54

#### InfluxDB价格
费用 = 时序节点费用 + 数据节点费用  

列如：以时序节点为3个，16核64GB规格；数据节点为9个，购买时长为1个月。则计费如下：  
时序节点费用 = 3 * 4160 = 12480元
数据节点费用 = 9 * 3000 = 27000元
实例总费用 = 12480 + 27000 = 39480元

### 附件
* [系统架构](https://cloud.tencent.com/document/product/652/89876)
* [产品定价](https://cloud.tencent.com/document/product/652/103982)

