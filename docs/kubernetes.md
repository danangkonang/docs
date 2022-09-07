# Belajar Kubernetes

## install
- v1.26.1 [minikube-linux-amd64.tar.gz](https://github.com/kubernetes/minikube/releases/download/v1.26.1/minikube-linux-amd64.tar.gz)

```bash
tar -xvf minikube-linux-amd64.tar.gz
```

```bash
#setelah install
minikube start --driver=docker
minikube config set driver docker

minikube start
minikube stop
```

```bash
kubectl version --server
kubectl version --client
kubectl get node
kubectl get pod
kubectl get all
kubectl get deployments
kubectl apply -f deployment.yml

minikube service golang-helloword-service
```