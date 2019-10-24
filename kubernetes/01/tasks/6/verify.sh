#!/bin/bash

[ `kubectl get deployment -n kube-system kubernetes-dashboard -o jsonpath='{.spec.template.spec.containers[0].name}'` == "kubernetes-dashboard" ] &&
[ `kubectl get svc -n kube-system kubernetes-dashboard -o jsonpath='{.spec.ports[0].targetPort}'` == "8443" ] &&
[ `kubectl get ep -n kube-system kubernetes-dashboard -o jsonpath="{.subsets[*].addresses[0].nodeName}"` == 'node01' ] &&
echo done