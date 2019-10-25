#!/bin/bash

[[ $(kubectl get deployments -n ingress-nginx nginx-ingress-controller -o jsonpath='{.status.readyReplicas}') == '1' ]] &&
echo done