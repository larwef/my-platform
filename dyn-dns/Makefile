include config.env

APP=dyn-dns
BUILD_VERSION=0.0.1
TARGET=target

# ------------------------------------- Go -------------------------------------
.PHONY: build
build: build-app docker-build

build-app:
	GOOS=linux GOARCH=arm GOARM=7 go build -ldflags "-X main.version=v$(BUILD_VERSION)" -o $(TARGET)/$(APP).bin cmd/$(APP)/main.go

.PHONY: test
test:
	go test ./...

coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out

# ---------------------------------- Docker -----------------------------------
.PHONY: docker
docker: docker-build docker-push

docker-build:
	docker build -t $(REPOSITORY)/$(APP):$(BUILD_VERSION) --build-arg target=$(TARGET)/$(APP).bin -f build/package/Dockerfile .

docker-push:
	docker push $(REPOSITORY)/$(APP):$(BUILD_VERSION)
