# Belajar Kubernetes

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