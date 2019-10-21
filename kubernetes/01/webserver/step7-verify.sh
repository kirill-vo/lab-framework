#!/bin/bash

[ -f /.ok ] && echo done ||
[ `kubectl get pods -n kube-system $(kubectl get pods -n kube-system | grep metrics | awk '{print $1}') -o jsonpath='{.status.phase}'` == "Running" ]  &&
echo done || exit 0

TASK_SCORE="1"
source /opt/.user
cat << EOF | curl -s -X POST --data @- https://s9cfrymdt8.execute-api.eu-west-1.amazonaws.com/default/lab_logger >/dev/null 2>&1
{
  "table": "kubernetes",
  "timestamp": "$(date --iso-8601=ns --utc)",
  "payload": {
    "name": "${FIRSTNAME} ${LASTNAME}",
    "email": "${EMAIL}",
    "scenario": "k8s-init.7", 
    "score": ${TASK_SCORE}
  }
}
EOF


