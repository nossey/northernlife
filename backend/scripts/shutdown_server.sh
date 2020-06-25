#!/bin/bash
PID=`lsof -t -i :9000`
if [ -n "$PID" ]; then
  kill $PID
fi
