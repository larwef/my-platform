include config.env

BUILD_VERSION=$(shell git describe --tag --always | sed -e 's,^v,,')

APP=hello-world

.PHONY: build
build:
	make docker-build

docker-build:
	docker build -t $(REPOSITORY)/$(APP):$(BUILD_VERSION) \
		--build-arg app_name=$(APP) \
		--build-arg build_verion=$(BUILD_VERSION) \
		--build-arg port=$(PORT) \
		-f build/package/Dockerfile .

compose-up:
	REPOSITORY=${REPOSITORY} \
	VERSION=${BUILD_VERSION} \
	DOMAIN=${DOMAIN} \
		docker compose -f deployments/docker-compose/docker-compose.yml up --remove-orphans --detach

compose-down:
	REPOSITORY=${REPOSITORY} \
	VERSION=${BUILD_VERSION} \
	DOMAIN=${DOMAIN} \
		docker compose -f deployments/docker-compose/docker-compose.yml down

compose-logs:
	(cd deployments/docker-compose && \
		REPOSITORY=${REPOSITORY} \
		VERSION=${BUILD_VERSION} \
		DOMAIN=${DOMAIN} \
		docker compose logs -f -t)
