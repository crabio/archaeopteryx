# archaeopteryx example

This is an example of usage `archaeopteryx` for creating microservice on Golang with all required dependencies

## Install dependencies

```sh
make install
```

## Generate protobuf & docs

```sh
make generate
```

## Presequinces

`archaeopteryx` requires folder for log files.
To create it run:
```sh
sudo mkdir /var/log/archaeopteryx
sudo chown $USER /var/log/archaeopteryx
```

## Build

### Docker

To build docker image use:
```sh
make build_docker
```

## Run

### Binary

To run server use:
```sh
make run
```

### Docker

To run server in docker use:
```sh
make run_docker
```

## Test

### Unit test

For unit tests use:
```sh
make test
```

### Lint

For lint use:
```sh
make lint
```

### gRPC

#### gRCP Curl

For getting info about method use:
```sh
docker run --network="host" fullstorydev/grpcurl -plaintext localhost:8080 describe grpc.health.v1.Health.Check
```

#### Kreya

For testing gRPC API use [Kreya](https://kreya.app/)

Folder `kreya` contains Kreya project for working with the project.