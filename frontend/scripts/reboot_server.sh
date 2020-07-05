#!/bin/bash

docker pull nossey/northernlife-frontend:latest
docker run -p 8000:3000 -d nossey/northernlife-frontend:latest
sudo service nginx start