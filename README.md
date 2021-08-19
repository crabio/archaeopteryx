# archaeopteryx
This project is a prototype for microservice on Golang with all required dependencies

## Structure

* `buf.yaml` - general configuration file for the `buf` generator with paths to `.proto` files
* `buf.gen.yaml` - configuration file for the `buf` generator for generating different stubs (`gRPC`, `swagger`, `REST`)

## Install dependencies

```sh
make install
```

## Generate protobuf & docs

```sh
make generate
```

## Run

To run server use
```sh
go run .
```

## Test

### gRPC gateway

To test REST gRPC gateway use command:
```sh
curl -X POST -k http://localhost:8090/v1/example/echo -d '{"name": " hello"}'
```

As a result you should get response: `{"message":" hello world"}`