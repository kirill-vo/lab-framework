#!/bin/bash

export KUBECONFIG=/etc/kubernetes/admin.conf

kubectl get nodes >/dev/null 2>&1 &&
kubeadm token list | grep abcdef.0123456789abcdef >/dev/null &&
[ `kubeadm token list | grep abcdef.0123456789abcdef | awk '{print $2}' | egrep -c '[01][0-9]m'` -eq 1 ] &&
echo done