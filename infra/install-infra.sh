#!/bin/bash
helm repo add stable https://kubernetes-charts.storage.googleapis.com
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm repo update
helm upgrade -i ingress ingress-nginx/ingress-nginx

helm upgrade -i -f registry.yaml docker-registry stable/docker-registry