#!/bin/bash

set -a
source .env
set +a
echo $DOCKER_TAG

# docker run -p 80:8000 -d nossey/northernlife-frontend:latest > /dev/null 2>&1 &