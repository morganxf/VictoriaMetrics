
### VMAgent CR增加volume

在spec下增加volumes和volumeMounts两个字段，字段内容与Deployment一样。
其中，secret需要mount到/etc/vm/secrets目录下

```yaml
spec:
  volumeMounts:
  - mountPath: /etc/vm/secrets/kube-etcd-client-certs
    name: secret-kube-etcd-client-certs
    readOnly: true
  volumes:
  - name: secret-kube-etcd-client-certs
    secret:
      defaultMode: 420
      secretName: kube-etcd-client-certs
```
