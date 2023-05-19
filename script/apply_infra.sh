#!/bin/sh

LAYER=$1
ENV=$2

cd tffile/environment/$LAYER/$ENV
terraform init -input=false -no-color
terraform apply -input=false -no-color -auto-approve
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml
sleep 3
kubectl wait --namespace ingress-nginx \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/component=controller \
  --timeout=90s
kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml
kubectl patch deployment metrics-server -n kube-system -p '{"spec":{"template":{"spec":{"containers":[{"name":"metrics-server","args":["--cert-dir=/tmp","--secure-port=4443","--v=2","--kubelet-insecure-tls","--kubelet-preferred-address-types=InternalIP"]}]}}}}'
sleep 10
kubectl wait --namespace kube-system \
  --for=condition=ready pod \
  --selector=k8s-app=metrics-server \
  --timeout=90s
