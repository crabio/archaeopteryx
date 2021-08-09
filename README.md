# go-microservice
This project is a prototype for microservice on Golang with all required dependencies


## Install Buf

```sh
GO111MODULE=on go get \
  github.com/bufbuild/buf/cmd/buf \
  github.com/bufbuild/buf/cmd/protoc-gen-buf-breaking \
  github.com/bufbuild/buf/cmd/protoc-gen-buf-lint
```

## Generate protobuf files

```sh
protoc -I .\
    --go_out ./proto/go\
    --go_opt paths=source_relative\
    --go-grpc_out ./proto/grpc\
    --go-grpc_opt paths=source_relative\
    ./proto/api.proto
```