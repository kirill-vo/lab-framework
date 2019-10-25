#!/bin/bash

kubectl get deployment -n kube-system metrics-server -o jsonpath='{.status.conditions[?(@.type=="Progressing")].status}' | grep True &&
kubectl top pods &&
kubectl top nodes