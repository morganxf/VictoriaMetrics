### 使用RBAC获取Metrics
```sh
# 指定vmagent namespace
namespace={vmagentNamespace}
# 指定vmagent pod name
pod={vmagentPodName}
# 指定 <host>:<port>
instance={host:port}
# 获取SA信息
mkdir vmagent;
kubectl -n ${namespace} exec ${pod} cat /var/run/secrets/kubernetes.io/serviceaccount/ca.crt > vmagent/ca.crt;
kubectl -n ${namespace} exec ${pod} cat /var/run/secrets/kubernetes.io/serviceaccount/token > vmagent/token;
kubectl -n ${namespace} exec ${pod} cat /var/run/secrets/kubernetes.io/serviceaccount/namespace > vmagent/namespace;
token=$(cat vmagent/token)
# 获取Metrics
curl --insecure --cacert vmagent/ca.crt --header "Authorization: Bearer ${token}" -X GET https://${instance}/metrics
```
#### 结果分析
* 成功：Exporter没有问题
* 失败：Exporter有问题
