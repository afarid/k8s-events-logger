# k8s-events-logger
Controller to log all kubernetes events to stdout of single pods, then it can be shipped to centralized storage by log aggregator


# Installation
- Change namespace and image name in [Makefile](./Makefile)
```makefile
export IMAGE := "afarid/k8s-events-logger:0.0.1"
export NAMESPACE := "kube-system"
```

# Deploy to current k8s context
```shell
# Build the docker image for k8s-events-logger
make build
# If you logged to remove docker registry
make push  
# If you working on local kind cluster
make load 
# Deploy to current k8s context
make deploy
```
