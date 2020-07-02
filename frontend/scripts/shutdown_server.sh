#!/bin/bash

PID=`sudo lsof -t -i :8000`

if [ -n "$PID" ]; then
  sudo kill $PID
fi