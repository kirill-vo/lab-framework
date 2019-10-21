
Deploy Dashboard Manifest from [here](https://raw.githubusercontent.com/kubernetes/dashboard/v1.10.1/src/deploy/recommended/kubernetes-dashboard.yaml)

## Create Admin User

```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: admin-user
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: admin-user
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
- kind: ServiceAccount
  name: admin-user
  namespace: kube-system
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
