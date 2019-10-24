#!/bin/bash

[[ $(kubectl get nodes worker | awk 'FNR==2{print $2}') == 'Ready' ]] &&
echo done