#!/bin/bash
set -e
export REGISTRY=docker-registry.default
export TAG="epoch-$(date +%s)"
echo "Registry: $REGISTRY"
echo "Tag: $TAG"

docker-compose -f docker-compose.services.yaml build
docker-compose -f docker-compose.yaml -f docker-compose.services.yaml up -d
docker-compose -f docker-compose.services.yaml run --rm integration-test  bash -c "go test -v"
docker-compose -f docker-compose.services.yaml push

pushd deploy
./deploy.sh
popd

sleep 5s
./end-to-end.sh
