# rpi-setup
A place to documenting basic setup of Raspberry Pi(s).

## Install Podman
https://podman.io/getting-started/installation#debian
```
sudo apt-get -y install podman
```
To avoid the warnings about cgroupv2 (reboot after command):
```
loginctl enable-linger <username>
```
## Running Container Using systemd and Podman in user_mode
https://podman.io/blogs/2020/12/09/podman-systemd-demo.html
https://github.com/edhaynes/podman_systemd_usermode_demo

## Useful Stuff
Reload systemd daemon:
```
sudo systemctl daemon-reload
```
Check service status:
```
sudo systemctl status echo-server.service
```

See service logs:
```
journalctl -f -u echo-server.service
```

See rpi temperature:
```
vcgencmd measure_temp
```