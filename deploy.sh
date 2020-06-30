#!/bin/bash
set -e
export REGISTRY=docker-registry.default
export TAG="epoch-$(date +%s)"
echo "Registry: $REGISTRY"
echo "Tag: $TAG"

docker-compose -f docker-compose.services.yaml build
docker-compose -f docker-compose.services.yaml push

pushd deploy
./deploy.sh
popd