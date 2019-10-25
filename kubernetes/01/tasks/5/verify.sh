#!/bin/bash

kubectl get pods -n kube-system | grep kindnet | grep Running >/dev/null && 
kubectl get nodes master -o jsonpath='{.status.conditions[?(@.type=="Ready")].status}' | grep True