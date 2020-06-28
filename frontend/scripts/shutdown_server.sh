#!/bin/bash

PID=`sudo lsof -t -i :8080`

if [ -n "$PID" ]; then
  sudo kill $PID
fi