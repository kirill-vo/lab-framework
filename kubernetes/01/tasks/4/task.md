# 4. Configuring `kubectl` bash Completion

Configure `kubectl` bash completion for user `user01` permanently

## Verification:
```
user01@master:/# kubectl <tab><tab>
annotate       autoscale      cordon         drain          kustomize      port-forward   set            
api-resources  certificate    cp             edit           label          proxy          taint          
api-versions   cluster-info   create         exec           logs           replace        top            
apply          completion     delete         explain        options        rollout        uncordon       
attach         config         describe       expose         patch          run            version        
auth           convert        diff           get            plugin         scale          wait  
```

## Documentation:
- https://kubernetes.io/docs/reference/kubectl/cheatsheet/#kubectl-autocomplete
