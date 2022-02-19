#!/bin/bash

# exit when any command fails
set -e

source ./variables.sh

podman create \
    --name $SERVICE \
    -p $HOST_PORT:$CONTAINER_PORT $IMAGE
podman generate systemd $SERVICE --restart-policy=always -t 5 -f -n
mkdir -p ~/.config/systemd/user
cp ./container-$SERVICE.service ~/.config/systemd/user/$SERVICE.service
systemctl enable --user $SERVICE.service
systemctl start --user $SERVICE.service
# systemctl status --user $SERVICE.service
