#!/bin/bash


kubectl get nodes worker -o json | jq '.metadata.labels."node-role.kubernetes.io/worker-node"' | grep '""'