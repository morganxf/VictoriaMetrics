
retention period. 怎么删除老数据 删除逻辑？
big merge ?
small merge ?
final merge ?
vminsterAddr + vmselectAddr 读写分离 端口也分离么？
snapshot 作用 ？
retentionTimezoneOffset 默认是UCT时间 4am，中国时区需要设置，保证在中国时区的 4am执行。

采集间隔保持简单，最好可以只有一个值。https://www.robustperception.io/keep-it-simple-scrape_interval-id/

deduplication 是否应该开启

indexdb/indexBlocks indexdb/dataBlocks indexdb/tagFiltersToMetricIDs
storage/tsid

vmstorage启动从磁盘加载数据流程，磁盘数据量大会导致启动慢么？如何优化？

8482端口：
force merge
force flush
snapshot create list delete delete_all
snapshot list
snapshot 

