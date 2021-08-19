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