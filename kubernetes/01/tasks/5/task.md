# 5. Deploying POD Network Driver

![network](https://assets.digitalocean.com/articles/k8s-networking/double-service.png)


## Requirements:
- Pod Network Manifest File is located here: `/opt/manifests/pod-network.yaml`


## Verification:
```
kubectl get pods -n kube-system | grep kindnet
kindnet-2rs42    1/1     Running   0          3m18s

kubectl get nodes
NAME     STATUS   ROLES    AGE     VERSION
master   Ready    master   7m35s   v1.15.3
```

## Tips:
- Check that CNI pods are running