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
curl -X POST -k http://localhost:8090/v2/user -d '{"fisrt_name":"bobby","last_name":"twist","password":"ilu"}'
```

As a result you should get response: `{"id":<random id>}`

##### Get User

To test REST gRPC gateway use command:
```sh
curl -X GET -k http://localhost:8090/v1/user?id=123
```

As a result you should get response: `{"name":"Bobby", "password":"qwerty"}`