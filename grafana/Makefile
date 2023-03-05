include config.env

compose-up:
	GRAFANA_VERSION=${GRAFANA_VERSION} \
	DOMAIN=${DOMAIN} \
		docker compose up --remove-orphans --detach

compose-down:
	GRAFANA_VERSION=${GRAFANA_VERSION} \
	DOMAIN=${DOMAIN} \
		docker compose down

compose-logs:
	GRAFANA_VERSION=${GRAFANA_VERSION} \
	DOMAIN=${DOMAIN} \
		docker compose logs -f -t