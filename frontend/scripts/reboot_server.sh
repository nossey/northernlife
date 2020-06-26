#!/bin/bash

sudo set -a
sudo source .env
sudo set +a

sudo docker pull nossey/northernlife-frontend:$DOCKER_TAG
sudo docker run -p 80:8000 -d nossey/northernlife-frontend:$DOCKER_TAG > /dev/null 2>&1 &