# Docker Registry

Create a docker registry inside k8s

## Setup k8s

Create the namespace and PersistentVolumeClaim
```
kubectl apply -f ./k8s
```

Add the twuni repo to your local helm

```
helm repo add twuni https://helm.twun.io
```
```
helm repo update
```

Install the registry

```
helm install docker-registry twuni/docker-registry -n docker-registry -f docker-registry-values.yaml
```

This will create the docker-registry.docker-registry service on port 5000 and create a nodePort on 30100.
This is required since the kubernetes nodes will need to be able to get to the registry using localhost:30100.

Using codezero consume the registry

```
czctl compose start -f docker-registry-compose.yaml
```

Test that it works

```
curl -I http://docker-registry.docker-registry:5000/
```

## Using the registry

If you local architecture is the same as your kubernetes architecture then you can use the following,
if not skip to the building in kubernetes section.

Build hello

```
docker build -t hello:latest hello
```

Tag it for the docker-registry

```
docker tag hello:latest docker-registry.docker-registry:5000/hello:latest
```

Push it to the docker-registry

```
docker push docker-registry.docker-registry:5000/hello:latest
```

Deploy hello

```
kubectl apply -f hello.yaml
```

Consume the hello service

```
czctl compose start -f hello-compose.yaml
```

Test

```
curl hello.hello-demo:8080
```

## Building inside kubernetes

Create a buildx builder to use the kubernetes cluster

```
kubectl create namespace buildkit
docker buildx create  \
  --config ./buildkitd.toml \
  --bootstrap \
  --name=kube \
  --driver=kubernetes \
  --driver-opt=namespace=buildkit,replicas=4,qemu.install=true
```

Make sure the deployments are running

```
kubectl -n buildkit get deployments
```

Build the docker images and push to the docker-registry

```
docker buildx build \
  --builder=kube \
  --platform=linux/amd64,linux/arm64 \
  -t docker-registry.docker-registry:5000/hello:latest \
  --push hello
```

Deploy hello

```
kubectl apply -f hello.yaml
```

Consume the hello service

```
czctl compose start -f hello-compose.yaml
```

Test

```
curl hello.hello-demo:8080
```

