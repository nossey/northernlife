#!/bin/bash
docker image pull nossey/northernlife-frontend:latest
docker image pull nossey/northernlife-backend:latest
cd /opt/app/
source ~/.bash_profile
docker-compose up