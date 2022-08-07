#!/bin/bash

# exit when any command fails
set -e

SERVICE_NAME=influxdb
PORT=8086

mkdir -p ~/data/influxdb
podman create \
    --name $SERVICE_NAME \
    --env-file $PWD/influxdb/envfile.env \
    -v ~/data/influxdb:/var/lib/influxdb2 \
    -p $PORT:$PORT docker.io/library/influxdb:2.3-alpine

setup-service/launch.sh $SERVICE_NAME
