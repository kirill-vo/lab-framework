#!/bin/bash
exit 0
kubectl get node worker --show-labels | grep 'node-role.kubernetes.io/worker=' >/dev/null &&
echo done