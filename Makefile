generate:
	echo "Generate proto"
	go mod tidy
	buf mod update
	buf generate

lint:
	echo "Run linter"
	# Lint proto
	buf lint
	# Lint golang
	golangci-lint run

test:
	echo "Run unit tests"
	go test -v ./...

run:
	echo "Run app"
	go run .

BUF_VERSION:=0.48.2

install:
	echo "Install Buf generators"
	go install \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	curl -sSL \
    	"https://github.com/bufbuild/buf/releases/download/v${BUF_VERSION}/buf-$(shell uname -s)-$(shell uname -m)" \
    	-o "$(shell go env GOPATH)/bin/buf" && \
  	chmod +x "$(shell go env GOPATH)/bin/buf"

	echo "Install linter"
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.42.0