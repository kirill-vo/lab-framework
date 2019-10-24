#!/bin/bash

kubectl get nodes worker -o jsonpath='{.status.conditions[?(@.type=="Ready")].status}' | grep True