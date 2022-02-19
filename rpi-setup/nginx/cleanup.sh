#!/bin/bash

# exit when any command fails
set -e

source ./variables.sh

systemctl --user stop $SERVICE.service
systemctl --user disable $SERVICE.service
rm -f ~/.config/systemd/user/$SERVICE.service
podman stop $SERVICE
podman rm $SERVICE
