#!/bin/bash

# exit when any command fails
set -e

SERVICE_NAME=postgres
IMAGE=docker.io/library/postgres:14-alpine
DATA_DIR=/var/lib/postgresql/data
PORT=5432

mkdir -p ~/data/$SERVICE_NAME
podman create \
    --name $SERVICE_NAME \
    --env-file $PWD/$SERVICE_NAME/envfile.env \
    --network=host \
    -v ~/data/$SERVICE_NAME:$DATA_DIR \
    -p $PORT:$PORT $IMAGE

setup-service/launch.sh $SERVICE_NAME
