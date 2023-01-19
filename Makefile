GOOS=linux
GOARCH=amd64
CGO_ENABLED=1


.PHONY: build
build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=$(CGO_ENABLED) \
	go build -o main cmd/main.go
