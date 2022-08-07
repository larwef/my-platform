# Run Influxdb as Container Managed by Systemd
`./launch.sh` will create a new container running Influxdb, generate systemd
service and start the service so it will run on boot.

A volume is mounted to make sure data is persisted even if the container is
stopped and/or removed.

To remove container and service run `./cleanup.sh`.

Remember to make a `envfile.env` and add env variables for the container. See
example.