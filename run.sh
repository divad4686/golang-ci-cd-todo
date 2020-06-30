#!/bin/bash
set -e
docker-compose -f docker-compose.services.yaml build

docker-compose -f docker-compose.yaml -f docker-compose.services.yaml up -d

