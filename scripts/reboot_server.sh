#!/bin/bash
docker pull nossey/northernlife-frontend:latest
docker pull nossey/northernlife-backend:latest
cd /var/application
docker-compose up