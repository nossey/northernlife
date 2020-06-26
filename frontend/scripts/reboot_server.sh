#!/bin/bash

set -a
source .env
set +a

docker pull nossey/northernlife-frontend:$DOCKER_TAG
docker run -p 80:8000 -d nossey/northernlife-frontend:$DOCKER_TAG > /dev/null 2>&1 &