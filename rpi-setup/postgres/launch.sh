#!/bin/bash

# exit when any command fails
set -e

SERVICE_NAME=postgres
PORT=5432

mkdir -p ~/data/postgres
podman create \
    --name $SERVICE_NAME \
    --env-file $PWD/postgres/envfile.env \
    -v ~/data/postgres:/var/lib/postgresql/data \
    -p $PORT:$PORT docker.io/library/postgres:14-alpine

setup-service/launch.sh $SERVICE_NAME
