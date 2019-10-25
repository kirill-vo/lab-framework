#!/bin/bash

mkdir -p /opt/manifests/
wget -O- https://raw.githubusercontent.com/kubernetes/kops/master/addons/metrics-server/v1.8.x.yaml | sed 's@apiVersion: extensions/v1beta1@apiVersion: apps/v1@' > /opt/manifests/metrics-server.yaml