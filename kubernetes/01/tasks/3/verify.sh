#!/bin/bash

kubectl get pods -n kube-system | grep weave | grep Running >/dev/null &&
echo done