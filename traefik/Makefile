compose-up:
	TRAEFIK_CONFIG=./traefik-local.toml \
		docker compose up --remove-orphans --detach

compose-down:
	TRAEFIK_CONFIG=./traefik-local.toml \
		docker compose down

compose-logs:
	docker compose logs -f -t