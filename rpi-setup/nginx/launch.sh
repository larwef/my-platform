#!/bin/bash

# exit when any command fails
set -e

SERVICE_NAME=nginx
HOST_PORT=8080
CONTAINER_PORT=80

podman create \
    --name $SERVICE_NAME \
    -p $HOST_PORT:$CONTAINER_PORT docker.io/library/nginx:1.21.6-alpine

setup-service/launch.sh $SERVICE_NAME
