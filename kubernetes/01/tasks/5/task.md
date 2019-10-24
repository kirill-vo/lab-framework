# Task 5

Label *worker* node as `worker`

```
kubectl get nodes
NAME     STATUS   ROLES    AGE     VERSION
master   Ready    master   3m43s   v1.14.0
worker   Ready    worker   46s     v1.14.0
```


## Tips:
- You should just label necessary node as `node-role.kubernetes.io/<< node role >>`

## Documentation:
- http://kubernetesbyexample.com/labels/
- https://kubernetes.io/docs/concepts/cluster-administration/manage-deployment/#updating-labels
