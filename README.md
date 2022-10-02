# parthenon

## cluster setup

### install kubectl

kubectl is a command line client for k8. [install page](https://kubernetes.io/docs/tasks/tools/)

### start k8 cluster

In docker click on settings (the gear icon), go to the kubernetes tab, and choose enable kubernetes. Wait for the cluster to startup. Run `kubectl get pod -A`, you should see system pods.

### setup k8 cluster

run `make emissary`, then `make apply`

### start dev server

run `yarn start`
