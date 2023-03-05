include config.env

compose-up:
	GRAFANA_VERSION=${GRAFANA_VERSION} \
		docker compose up --remove-orphans --detach

compose-down:
	GRAFANA_VERSION=${GRAFANA_VERSION} \
		docker compose down

compose-logs:
	GRAFANA_VERSION=${GRAFANA_VERSION} \
		docker compose logs -f -t