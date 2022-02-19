#!/bin/bash

# exit when any command fails
set -e

source ./variables.sh

mkdir -p ~/data/postgres
podman create \
    --name $SERVICE \
    --env-file ./envfile.env \
    -v ~/data/postgres:/var/lib/postgresql/data \
    -p $PORT:$PORT $IMAGE
podman generate systemd $SERVICE --restart-policy=always -t 5 -f -n
mkdir -p ~/.config/systemd/user
cp ./container-$SERVICE.service ~/.config/systemd/user/$SERVICE.service
systemctl enable --user $SERVICE.service
systemctl start --user $SERVICE.service
# systemctl status --user $SERVICE.service
