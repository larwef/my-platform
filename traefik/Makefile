include config.env

compose-up:
	TRAEFIK_VERSION=${TRAEFIK_VERSION} \
		docker compose up --remove-orphans --detach

compose-down:
	TRAEFIK_VERSION=${TRAEFIK_VERSION} \
		docker compose down

compose-logs:
	TRAEFIK_VERSION=${TRAEFIK_VERSION} \
		docker compose logs -f -t