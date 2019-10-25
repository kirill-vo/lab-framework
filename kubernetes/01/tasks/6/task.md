# 5. Labeling Resources

## Requirements:
- Label *worker* node as `worker-node`

```
kubectl get nodes
NAME     STATUS   ROLES         AGE     VERSION
master   Ready    master        3m43s   v1.15.3
worker   Ready    worker-node   46s     v1.15.3
```

## Tips:
- You should just label necessary node as `node-role.kubernetes.io/<< node role >>`
- Use commande like `kubectl label node <nodename> <labelname>=<labelvalue>`
- If you need to label resource with empty label just ommit `<labelvalue>`

## Documentation:
- http://kubernetesbyexample.com/labels/
- https://kubernetes.io/docs/concepts/cluster-administration/manage-deployment/#updating-labels
