# archaeopteryx

[![codecov](https://codecov.io/gh/iakrevetkho/archaeopteryx/branch/develop/graph/badge.svg?token=3QSO8BJHJA)](https://codecov.io/gh/iakrevetkho/archaeopteryx)

This project is a prototype for microservice on Golang with all required dependencies

## How to use

`archaeopteryx` helps you to create service with gRPC, gRPC proxy for REST API and some editional features from the box.
All features can be found in the list

### Config

Service has configuration, bases on https://github.com/jinzhu/configor.
Your application should reimplement config of server like:
```go
import (
	archaeopteryx_config "github.com/iakrevetkho/archaeopteryx/pkg/config"
)

type Config struct {
	archaeopteryx_config.Config
}
```

### Logger

Service has methods to create logger:
```go
import (
	"github.com/sirupsen/logrus"
	archaeopteryx_config "github.com/iakrevetkho/archaeopteryx/logger"
)

func main(){
    var log *logrus.Entry
    log = logger.CreateLogger("main")
}
```

All messages will be formatted with component name to make easy search by this name:

![Logs example](docs/img/logs.png)

### Service server interface

All custom services should implements interface `service_server`:
```go
// IServiceServer - interface for services servers for archaeopteryx.
//
// All services should implements this interface
// to make possible archaeopteryx registrate services handlers in server
type IServiceServer interface {
	// RegisterGrpc - HealthServiceServer's method to registrate gRPC service server handlers
	RegisterGrpc(sr grpc.ServiceRegistrar) error
	// RegisterGrpcProxy - HealthServiceServer's method to registrate gRPC proxy service server handlers
	RegisterGrpcProxy(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
}
```

## Features

* health check bases on https://github.com/grpc/grpc/blob/master/doc/health-checking.md
* logger with logrus nested formatter (https://github.com/antonfisher/nested-logrus-formatter)
* logger writes data into the console and file, which is rotated with https://github.com/iakrevetkho/woodpecker
* config is inited bases on https://github.com/jinzhu/configor
* server is waiting terminate signals on run

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