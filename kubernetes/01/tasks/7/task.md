# Task 7

Deploy Metrics Server

## Requirements:
- Metrics Server Manifest is located here: `/opt/metrics-server.yaml`

## Check that:
- Metrics Server is deployed successfully and is `Running`
- Command `kubectl top nodes` shows performance statistics by nodes
- Command `kubectl top pods --all-namespaces` shows performance statistics by pods

## Documentation:
- https://kubernetes.io/docs/tasks/debug-application-cluster/resource-usage-monitoring/
- https://github.com/kubernetes-incubator/metrics-server
