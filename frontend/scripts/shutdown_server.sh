#!/bin/bash

PID=`sudo lsof -t -i :8080`
echo $PID

#if [ -n "$PID" ]; then
#  kill $PID
#fi