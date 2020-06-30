#!/bin/bash
# REGISTRY=docker-registry.default
# TAG=epoch-1592389230

environment=${ENVIRONMENT:-'staging'}

chart='todo-cicd'
deploy=$chart-$environment
namespace=$chart-$environment

kubectl create namespace $namespace
helm dep update $chart


helm upgrade --install \
  $deploy \
  --namespace $namespace \
  --set imageRegistry=${REGISTRY} \
  --set imageTag=${TAG} \
  --set namespace=$namespace \
  --set environment=$environment \
  --debug \
  $chart \
  -f $chart/values.yaml \
  --wait
