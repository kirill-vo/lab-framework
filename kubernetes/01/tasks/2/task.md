# 2. Client Access to the Cluster - 1

## Requiremets:
- user `root` should be able to use tool kubectl for managing k8s cluster (localhost)

## Verification:
```
user01@master:~$ kubectl get componentstatuses
NAME                 STATUS    MESSAGE             ERROR
controller-manager   Healthy   ok                  
scheduler            Healthy   ok                  
etcd-0               Healthy   {"health":"true"} 
```

## Documentation:
- https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/create-cluster-kubeadm/
