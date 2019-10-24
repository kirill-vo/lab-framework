#!/bin/bash -x

cat << EOF | kubectl apply -f -
kind: ConfigMap
apiVersion: v1
metadata:
  name: coredns
  namespace: kube-system
data:
  Corefile: |
    .:53 {
        errors
        health
        ready
        kubernetes cluster.local in-addr.arpa ip6.arpa {
           pods insecure
           fallthrough in-addr.arpa ip6.arpa
           ttl 30
        }
        prometheus :9153
        forward . /etc/resolv.conf
        cache 30
        reload
        loadbalance
    }
EOF

kubectl patch deployments.apps -n kube-system coredns -p '{
  "spec": {
    "replicas": 1, 
    "template": {
      "spec": {
        "containters": [
          {
            "name": "coredns", 
            "nodeSelector": {
              "kubernetes.io/hostname": "master"
            }
          }
        ]
      }
    }
  }
}'

kubectl rollout restart deployment -n kube-system coredns