### SpringBoot Actuator Prometheus指标

| 指标名 | 描述 | 类型 | 标签 | 备注 |
| --- | --- | --- | --- | --- |
| jvm_gc_live_data_size_bytes | Size of long-lived heap memory pool after reclamation | gauge | | |
| jvm_gc_memory_promoted_bytes_total | Count of positive increases in the size of the old generation memory pool before GC to after GC | counter | | |
| jvm_buffer_memory_used_bytes | An estimate of the memory that the Java virtual machine is using for this buffer pool | gauge | id | |
| jvm_buffer_total_capacity_bytes | An estimate of the total capacity of the buffers in this pool | gauge | id | |
| jvm_memory_used_bytes | The amount of used memory | gauge | area,id | |
| jvm_gc_pause_seconds | Time spent in GC pause | summary | action,cause | _count,_sum,_max |
| jvm_threads_daemon_threads | The current number of live daemon threads | gauge | |
| jvm_classes_unloaded_classes_total | The total number of classes unloaded since the Java virtual machine has started execution | counter | | |
| jvm_threads_live_threads | The current number of live threads including both daemon and non-daemon threads | gauge | | |
| jvm_memory_max_bytes | The maximum amount of memory in bytes that can be used for memory | gauge | area,id | |
| jvm_gc_overhead_percent | An approximation of the percent of CPU time used by GC activities over the last lookback period or since monitoring began, whichever is shorter, in the range [0..1] | gauge | | |
| jvm_gc_max_data_size_bytes | Max size of long-lived heap memory pool | gauge | | |
| jvm_memory_usage_after_gc_percent | The percentage of long-lived heap pool used after the last GC event | gauge | area,pool | |
| jvm_buffer_count_buffers | An estimate of the number of buffers in the pool | gauge | id | |
| jvm_gc_memory_allocated_bytes_total | Incremented for an increase in the size of the (young) heap memory pool after one GC to before the next | counter | | |
| jvm_memory_committed_bytes | The amount of memory in bytes that is committed for the Java virtual machine to use | gauge | area,id | |
| jvm_classes_loaded_classes | The number of classes that are currently loaded in the Java virtual machine | gauge | | |
| jvm_threads_peak_threads | The peak live thread count since the Java virtual machine started or peak was reset | gauge | | |
| jvm_threads_states_threads | The current number of threads | gauge | state | |
| process_uptime_seconds | The uptime of the Java virtual machine | gauge | | |
| process_files_open_files | The open file descriptor count | gauge | | |
| process_start_time_seconds | Start time of the process since unix epoch | gauge | | |
| process_files_max_files | The maximum file descriptor count | gauge | | |
| process_cpu_usage | The "recent cpu usage" for the Java Virtual Machine process | gauge | | |
| tomcat_sessions_alive_max_seconds | 会话活跃的最长时间 |  gauge | | |
| tomcat_sessions_created_sessions_total | 累计创建的会话数 | counter | | |
| tomcat_sessions_expired_sessions_total | 累计过期的会话数 | counter | | |
| tomcat_sessions_active_current_sessions | 当前活跃的会话数 | gauge | | |
| tomcat_sessions_active_max_sessions | 最大活跃会话数 | gauge | | |
| tomcat_sessions_rejected_sessions_total | 累计拒绝的会话数 | counter | | |
| http_server_requests_seconds | Duration of HTTP server request handling | summary | exception,method,outcome,status,uri | _count,_sum,_max |
| disk_total_bytes | Total space for path | gauge | path | |
| disk_free_bytes | Usable space for path | gauge | path | |
| executor_queued_tasks | The approximate number of tasks that are queued for execution | gauge | name | |
| executor_pool_max_threads | The maximum allowed number of threads in the pool | gauge | name | |
| executor_queue_remaining_tasks | The number of additional elements that this queue can ideally accept without blocking | gauge | name | |
| executor_active_threads | The approximate number of threads that are actively executing tasks | gauge | name | |
| executor_pool_size_threads | The current number of threads in the pool | gauge | name | |
| executor_completed_tasks_total | The approximate total number of tasks that have completed execution | counter | name | |
| executor_pool_core_threads | The core number of threads for the pool | gauge | name | |
| cache_eviction_weight_total | The sum of weights of evicted entries. This total does not include manual invalidations | counter | cache,cacheManager,name | |
| cache_evictions_total | The number of times the cache was evicted | counter | cache,cacheManager,name | |
| cache_puts_total | The number of entries added to the cache | counter | cache,cacheManager,name | |
| cache_gets_total | The number of times cache lookup methods have returned a cached (hit) or uncached (newly loaded or null) value (miss) | counter | cache,cacheManager,name,result | |
| cache_size | The number of entries in this cache. This may be an approximation, depending on the type of cache | gauge | cache,cacheManager,name | |
| grpc_server_processing_duration_seconds | The total time taken for the server to complete the call | summary | method,methodType,service,statusCode | _sum,_count,_max |
| grpc_server_responses_sent_messages_total | The total number of responses sent | counter | method,methodType,service | |
| grpc_server_requests_received_messages_total | The total number of requests received | counter | method,methodType,service | |
| log4j2_events_total | Number of fatal level log events | counter | level | |
| application_ready_time_seconds | Time taken (ms) for the application to be ready to service requests | gauge | main_application_class | |
| application_started_time_seconds | Time taken (ms) to start the application | gauge | main_application_class | |
| system_cpu_usage | The "recent cpu usage" of the system the application is running in | gauge | | |
| system_cpu_count | The number of processors available to the Java virtual machine | gauge | | |
| system_load_average_1m | The sum of the number of runnable entities queued to available processors and the number of runnable entities running on the available processors averaged over a period of time | gauge | | |

### 参考文档
https://juejin.cn/post/7023321727107072013
https://blog.csdn.net/ssehs/article/details/123961221
https://iotdb.apache.org/zh/UserGuide/V0.13.x/Maintenance-Tools/Metric-Tool.html#_1-3-iotdb%E9%83%BD%E6%9C%89%E5%93%AA%E4%BA%9B%E7%9B%91%E6%8E%A7%E6%8C%87%E6%A0%87
