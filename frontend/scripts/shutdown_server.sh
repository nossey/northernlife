#!/bin/bash

sudo service nginx stop

#PID=`sudo lsof -t -i :8000`
#
#if [ -n "$PID" ]; then
#  sudo kill $PID
#fi

CONTAINER_ID=`sudo -S docker ps -q  --filter ancestor=nossey/northernlife-frontend:latest`
docker kill $CONTAINER_ID