# archaeopteryx_example

This project is a example of usage `archaeopteryx` for creating microservice on Golang with all required dependencies

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

For testing gRPC API use [Kreya](https://kreya.app/)

Folder `kreya` contains Kreya project for working with the project.

### gRPC gateway

#### Hello world

To test REST gRPC gateway use command:
```sh
curl -X POST -k http://localhost:8090/v1/hello -d '{"name": " hello"}'
```

As a result you should get response: `{"message":" hello world"}`

#### User V1

##### Add User

To test REST gRPC gateway use command:
```sh
curl -X POST -k http://localhost:8090/v1/user -d '{"name":"bobby","password":"ilu"}'
```

As a result you should get response: `{"id":<random id>}`

##### Get User

To test REST gRPC gateway use command:
```sh
curl -X GET -k http://localhost:8090/v1/user?id=123
```

As a result you should get response: `{"name":"Bobby", "password":"qwerty"}`

#### User V2

##### Add User

To test REST gRPC gateway use command:
```sh
curl -X POST -k http://localhost:8090/v2/user -d '{"firstName":"bobby","lastName":"twist","password":"ilu"}'
```

As a result you should get response: `{"id":<random id>}`

##### Get User

To test REST gRPC gateway use command:
```sh
curl -X GET -k http://localhost:8090/v2/user?id=123
```

As a result you should get response: `{"firstName":"Bobby", "lastName":"Twist", "password":"qwerty"}`