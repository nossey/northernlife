#!/bin/bash

sudo service nginx stop

CONTAINER_ID=`sudo -S docker ps -q  --filter ancestor=nossey/northernlife-frontend:latest`
docker kill $CONTAINER_ID