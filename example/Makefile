help:
	@echo "Makefile for the archaeopteryx_boilerplate."
	@echo "Available targets:"
	@echo " help - print help information"
	@echo " install - install required dependecies for the project"
	@echo " lint - run linter"
	@echo " test - run unit tests"
	@echo " build - build application binary"
	@echo " run - run application"
	@echo " build_docker - build application docker image"
	@echo " run_docker - run application in docker"

install:
	echo "Install linter"
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.42.0

lint:
	@echo "Run linter"
	# Lint proto
	buf lint
	# Lint golang
	golangci-lint run

test:
	@echo "Run unit tests"
	go test -v ./... -coverprofile coverage.txt -covermode atomic
	@echo "Code coverage"
	go tool cover -func coverage.txt

build:
	@echo "Build app binary"
	go build

run:
	@echo "Run app"
	go run .

build_docker:
	@echo "Build docker image"
	docker build . --file Dockerfile --tag archaeopteryx_boilerplate

run_docker:
	@echo "Run app in docker"
	docker run -p 8080:8080 -e LOG_LEVEL=trace  archaeopteryx_boilerplate
