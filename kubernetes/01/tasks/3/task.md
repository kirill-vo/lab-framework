# 3. Client Access to the Cluster - 2

We have just created a user `user01`. Please do necessary configuration for this user so that it will be able to manage kubernetes cluster with `kubectl` tool

## Requiremets:
- user `user01` should be able to use tool kubectl for managing k8s cluster (localhost)
- Use `sudo` to switch between user accounts
- User proper ownerships on required configuration file(s)

## Verification:
```
root@master:/# kubectl get componentstatuses 
NAME                 STATUS    MESSAGE             ERROR
controller-manager   Healthy   ok                  
scheduler            Healthy   ok                  
etcd-0               Healthy   {"health":"true"}
```

## Documentation:
- https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/create-cluster-kubeadm/
