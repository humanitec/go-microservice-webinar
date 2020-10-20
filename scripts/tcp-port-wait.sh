#!/bin/bash

if [ -z "$1" -o -z "$2" ]; then
  echo "tcp-port-wait.sh - block until specified TCP port becomes available"
  echo "Usage: tcp-port-wait.sh HOST PORT"
  exit 1
fi
echo Waiting for port $1:$2 to become available...
nc -z $1 $2
while ! nc -z $1 $2 2>/dev/null; do
  let elapsed=elapsed+1
  if [ "$elapsed" -gt 60 ]; then
    echo "TIMED OUT !"
    exit 1
  fi
  sleep 1
done

echo "READY !"
