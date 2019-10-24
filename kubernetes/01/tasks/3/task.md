# Task 3

Deploy POD Network Driver


## Requirements:
- Pod Network Manifest File is located here: `/opt/pod-network.yaml`


## Verification:
```
kubectl get pods -n kube-system | grep kindnet
kindnet-2rs42    1/1     Running   0          3m18s

kubectl get nodes
NAME     STATUS   ROLES    AGE     VERSION
master   Ready    master   7m35s   v1.16.2
```

## Tips:
- Check that CNI pods are running