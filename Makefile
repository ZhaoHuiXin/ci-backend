#!/bin/zsh
server:
	export GO111MODULE=on && export GOPROXY=https://goproxy.io && go mod download
	go build -tags netgo -a -o ./bin/backend main.go

build: server
	docker build -t $(IMAGE_URL):$(DOCKER_TAG) --cache-from $(IMAGE_URL):latest .

test:
	echo "Testing for project $(DOCKER_TAG)"