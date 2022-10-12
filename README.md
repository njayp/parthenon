# parthenon

## tools

### kubectl

brew install kubectl

### protoc

brew install protobuf

## cluster setup

### start k8 cluster

In docker click on settings (the gear icon), go to the kubernetes tab, and choose enable kubernetes. Wait for the cluster to startup. Run `kubectl get pod -A`, you should see system pods.

### setup k8 cluster

run `make emissary`, then `make build apply`

## ts dev server

`make start`
