#!/bin/bash

PID=`sudo lsof -t -i :8080`
echo $PID
echo $DOCKER_TAG

#if [ -n "$PID" ]; then
#  kill $PID
#fi