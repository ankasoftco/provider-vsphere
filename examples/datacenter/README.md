# Datacenter Example
First we need to create vSphere secret and provider.
```
kubectl create -f secret.yaml
kubectl create -f providerconfig.yaml
```
After that we can create 'Datacenter' resource via Crossplane.
```
kubectl create -f datacenter.yaml
```
