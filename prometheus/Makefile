include config.env

compose-up:
	PROMETHEUS_VERSION=${PROMETHEUS_VERSION} \
	DOMAIN=${DOMAIN} \
		docker compose up --remove-orphans --detach

compose-down:
	PROMETHEUS_VERSION=${PROMETHEUS_VERSION} \
	DOMAIN=${DOMAIN} \
		docker compose down

compose-logs:
	PROMETHEUS_VERSION=${PROMETHEUS_VERSION} \
	DOMAIN=${DOMAIN} \
		docker compose logs -f -t