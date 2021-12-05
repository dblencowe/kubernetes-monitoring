# Kubernetes Monitoring Tool
This proof of concept tool can be used for quickly tracking the
number of running pods within a Kubernetes cluster.

As it is proof of concept this tool currently reads your local
KUBECONFIG path from an environment variable and connects to that
cluster

## Installation
The tool is written in Golang. To compile and run the program first
install your Golang development environment and run the following commands

```shell
go build
KUBECONFIG=/Users/you/.kube/config ./kubernetes monitoring
```

## Configuration
|Variable| Values |Description|
|--------|--------|-----------|
|KUBECONFIG| <>     |The path to your KUBECONFIG file|
|OUTPUT_FORMAT| json / text | Output as JSON or Human Friendly|   