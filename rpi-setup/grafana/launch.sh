#!/bin/bash

# exit when any command fails
set -e

SERVICE_NAME=grafana
IMAGE=docker.io/grafana/grafana-oss:9.0.6
DATA_DIR=/var/lib/grafana
PORT=3000

mkdir -p ~/data/$SERVICE_NAME
# Need this or else Grafana can't write to the data directory.
chmod -R 777 ~/data/$SERVICE_NAME 
podman create \
    --name $SERVICE_NAME \
    --env-file $PWD/$SERVICE_NAME/envfile.env \
    -v ~/data/$SERVICE_NAME:$DATA_DIR \
    -p $PORT:$PORT $IMAGE

setup-service/launch.sh $SERVICE_NAME