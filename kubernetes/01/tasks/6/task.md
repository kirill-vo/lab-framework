# 6. Kubernetes Dashboard

![dashboard](https://github.com/kubernetes/dashboard/raw/master/docs/images/dashboard-ui.png)

## Requirements:
- Please use following manifest: [dashboard.yaml](https://raw.githubusercontent.com/kubernetes/dashboard/v1.10.1/src/deploy/recommended/kubernetes-dashboard.yaml)

## Create Admin User

```
kubectl create serviceaccount admin-user -n kube-system
kubectl create clusterrolebinding admin-user-clusterrolebinding \
  --serviceaccount=kube-system:admin-user \
  --clusterrole=cluster-admin
```

## Get Admin Secret

```
SECRET_NAME="$(kubectl -n kube-system get secret | grep admin-user | awk '{print $1}')"
kubectl get secret -n kube-system ${SECRET_NAME} -o jsonpath='{.data.token}' | base64 -d
```

## Proxy Dashboard

Expose Dashord from inside cluster with `kubectl proxy` command:

```
kubectl proxy --address='0.0.0.0' --accept-hosts='^*$'
```

## Explore Kubernetes Cluster on "Dashboard" Tab

Please use token authentication from previous step

## Documentation:
- https://github.com/kubernetes/dashboard
- https://github.com/kubernetes/dashboard/wiki/Creating-sample-user
