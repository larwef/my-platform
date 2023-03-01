REPOSITORY=github.com/larwef
BUILD_VERSION=$(shell git describe --tag --always | sed -e 's,^v,,')
TARGET=target

.PHONY: build
build:
	make hello-world

hello-world:
	make docker-build APP=hello-world BUILD_VERSION=$(BUILD_VERSION)

# ---------------------------------- Docker -----------------------------------
docker-build:
	docker build -t $(REPOSITORY)/$(APP):$(BUILD_VERSION) \
		--build-arg app_name=$(APP) \
		--build-arg build_verion=$(BUILD_VERSION) \
		-f build/package/Dockerfile .

compose-up:
	REPOSITORY=${REPOSITORY} \
	VERSION=${BUILD_VERSION} \
		docker compose -f deployments/docker-compose/docker-compose.yml up --remove-orphans --detach

compose-down:
	REPOSITORY=${REPOSITORY} \
	VERSION=${BUILD_VERSION} \
		docker compose -f deployments/docker-compose/docker-compose.yml down

compose-logs:
	(cd deployments/docker-compose && \
		REPOSITORY=${REPOSITORY} \
		VERSION=${BUILD_VERSION} \
		docker compose logs -f -t)
