GOOS=linux
GOARCH=amd64
CGO_ENABLED=1

DOCKER_IMAGE_REGISTRY=ghcr.io
DOCKER_IMAGE_OWNER=unitedcollab
DOCKER_IMAGE_NAME=pdf-preview
DOCKER_IMAGE_LABEL=latest

.PHONY: build
build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=$(CGO_ENABLED) \
	go build -o main cmd/main.go

.PHONY: build-docker-image
build-docker-image:
	docker build -t $(DOCKER_IMAGE_REGISTRY)/$(DOCKER_IMAGE_OWNER)/$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_LABEL) . --no-cache

.PHONY: push-docker-image
push-docker-image:
	docker push $(DOCKER_IMAGE_REGISTRY)/$(DOCKER_IMAGE_OWNER)/$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_LABEL)
	