# my-platform
Just experimenting.

## Setup
- Install docker
- Create network to be able to communicate between compose files: `docker network create common`
- Create `acme.json` file in the traefik folder and give it sufficient permissions: `touch traefik/acme.json && chmod 600 traefik/acme.json`
- Set up logrotate configuration for access log. This example will rotate each week and keep 12 weeks:
`sudo nano /etc/logrotate.d/apt`
```
<where you located the project>/my-platform/traefik/logs/access.log {
  rotate 12
  weekly
  compress
  missingok
  notifempty
}
```