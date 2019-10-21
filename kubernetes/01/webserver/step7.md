
Deploy Metrics Server from [here](https://raw.githubusercontent.com/kubernetes/kops/master/addons/metrics-server/v1.8.x.yaml)

## Check that:
- Metrics Server is deployed successfully and is `Running`
- Command `kubectl top nodes` shows performance statistics by nodes
- Command `kubectl top pods --all-namespaces` shows performance statistics by pods

## Documentation:
- https://kubernetes.io/docs/tasks/debug-application-cluster/resource-usage-monitoring/
- https://github.com/kubernetes-incubator/metrics-server
