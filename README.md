# Лабораторная работа №7

## Задание

```sh
$ make
# run gateway
$ make run-gateway
Server listening on 0.0.0.0:8080
# in the new window run grpc service
$ make run-grpc-service 
Server listening on 0.0.0.0:9090
```

```sh
# send request to gRPC server via gateway
$ curl --header "Content-Type: application/json" \
  --request POST \
  --data '
{
  "data": "world"
}
  ' \
  http://localhost:8080/v1/hello
{
  "data": "Hello world"
}
```

## Links

[gRPC gateway](https://github.com/grpc-ecosystem/grpc-gateway)
