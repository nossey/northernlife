#!/bin/bash

#sudo set -a
#sudo source .env
#sudo set +a

#echo $DOCKER_TAG

docker pull nossey/northernlife-frontend:latest
#docker pull nossey/northernlife-frontend:$DOCKER_TAG

#docker run -p 8000:3000 -d nossey/northernlife-frontend:$DOCKER_TAG > /dev/null 2>&1 &
docker run -p 8000:3000 -d nossey/northernlife-frontend:latest

sudo service nginx start