# Task 1


## Cluster Infrastructure

![cluster](https://mapr.com/blog/kubernetes-kafka-event-sourcing-architecture-patterns-and-use-case-examples/assets/clusters.png)


Initialize Master node with `kubeadm`

## Parameters:
- token: `abcdef.0123456789abcdef`
- token life duration: `20m`
- Pod Network CIDR: `10.244.0.0/16`

## Tips:
- To initialize cluster control plane, please run following command `kubeadm init`. Please look for required options
- You can destroy cluster configuration with `kubeadm reset cluster`
- Use `--ignore-preflight-errors=all` to disable SWAP and IPTABLES setttings which fail preflight checks

## Documentation:
- https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/create-cluster-kubeadm/
