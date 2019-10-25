# 9. Deploying Metrics Server

![metrics](https://user-images.githubusercontent.com/21168270/46579266-95846680-ca40-11e8-86d3-a42291476db8.png)

## Requirements:
- Metrics Server Manifest is located here: `/opt/manifests/metrics-server.yaml`
- If Kubernetes responds with `error: metrics not available yet` just wait for a while

## Check that:
- Metrics Server is deployed successfully and is `Running`
- Command `kubectl top nodes` shows performance statistics by nodes
- Command `kubectl top pods --all-namespaces` shows performance statistics by pods

## Documentation:
- [resource-usage-monitoring](https://kubernetes.io/docs/tasks/debug-application-cluster/resource-usage-monitoring/)
- [metrics-server](https://github.com/kubernetes-incubator/metrics-server)
