# gRPC + Gateway server app sample

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
