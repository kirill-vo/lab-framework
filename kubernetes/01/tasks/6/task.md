# 6. Joining `worker` Node to the cluster 

![worker](https://miro.medium.com/max/926/1*JZ8cm65P2_iLOIo051aWtQ.png)

## Tips:
- Remember token is `abcdef.0123456789abcdef`?
- Default API Port on master is `6443`
- Use `ssh worker` to connect to `worker` host
- Use also `--discovery-token-unsafe-skip-ca-verification` and  `--ignore-preflight-errors=all` options
- Wait till `worker` turns to `Ready` state

## Documentation:
- https://www.weave.works/blog/weave-net-kubernetes-integration/
- https://www.weave.works/docs/net/latest/kubernetes/kube-addon/
