#!/bin/bash

sudo service nginx stop

PID=`sudo lsof -t -i :8000`

if [ -n "$PID" ]; then
  sudo kill $PID
fi