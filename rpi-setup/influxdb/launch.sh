#!/bin/bash

# exit when any command fails
set -e

SERVICE_NAME=influxdb
IMAGE=docker.io/library/influxdb:2.3-alpine
DATA_DIR=/var/lib/influxdb2
PORT=8086

mkdir -p ~/data/$SERVICE_NAME
podman create \
    --name $SERVICE_NAME \
    --env-file $PWD/$SERVICE_NAME/envfile.env \
    --network=host \
    -v ~/data/$SERVICE_NAME:$DATA_DIR \
    -p $PORT:$PORT $IMAGE

setup-service/launch.sh $SERVICE_NAME
