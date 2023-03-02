# my-platform
Just experimenting.

## Setup
- Install docker
- Create network to be able to communicate between compose files: `docker network create common`
- Create `acme.json` file in the traefik folder and give it sufficient permissions: `touch traefik/acme.json && chmod 600 traefik/acme.json`