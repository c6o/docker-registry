# Docker Registry

Create a docker registry inside k8s

## Setup k8s

```
kubectl apply -f ./k8s
```

```
helm repo add twuni https://helm.twun.io
```

```
helm repo update

```


```
helm install docker-registry twuni/docker-registry -n docker-registry -f docker-registry-values.yaml
```

Consume the registry

```
czctl compose start -f docker-registry-compose.yaml
```

Test that it works

```
curl -I http://docker-registry.docker-registry:5000/
```
