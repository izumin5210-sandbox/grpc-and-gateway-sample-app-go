# gRPC + Gateway server app sample
[![Build Status](https://travis-ci.org/izumin5210-sandbox/grpc-and-gateway-sample-app-go.svg?branch=master)](https://travis-ci.org/izumin5210-sandbox/grpc-and-gateway-sample-app-go)v

## Getting Started
You can development this app using [creasty/rid](https://github.com/creasty/rid).

```
$ rid bootstrap
$ rid pg start
$ rid pg createdb grpc-and-gateway-sample-app-go
```

## Development
### Migration
TBD

### Run server
```
$ rid server
```

#### Sample request
```
$ curl --request GET \
  --url http://localhost:3100/api/profiles/1
```

### Run test
```
$ rid make test
```

### Install dependencies
```
$ rid make setup
```

### Add new dependency
```
$ rid dep -v -no-vendor
$ rid dep -v -add <pkg>
```

### Show JSON API document
```
$ rid swagger start
$ open http://localhost:8080
```
