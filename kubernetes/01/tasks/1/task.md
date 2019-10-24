# Task 1

Initialize Master node with `kubeadm`

## Parameters:
- token: `abcdef.0123456789abcdef`
- token life duration: `20m`

## Tips:
- To initialize cluster control plane, please run following command `kubeadm init`. Please look for required options
- You can destroy cluster configuration with `kubeadm reset cluster`
- Use `--ignore-preflight-errors=all` to disable SWAP and IPTABLES settings which fail preflight checks

## Documentation:
- https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/create-cluster-kubeadm/
