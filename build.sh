#!/bin/bash
set -e
export REGISTRY=docker-registry.default
export TAG="epoch-$(date +%s)"
echo "Registry: $REGISTRY"
echo "Tag: $TAG"

docker-compose build
docker-compose push

# pushd deploy
# ./deploy.sh
# popd