# Лабораторная работа №7

## Иллюстрация

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

## Задание

0. Создать директорию `go/src/github.com/<group-repo>/` и </br>
склонировать в нее данный репозиторий, изменив в файле **gateway.go**</br>
`github.com/bmstu-iu8-cpp-sem-3/lab-07-grpc/pkg/echo` на </br>
`github.com/<group-repo>/<repo-name>/pkg/echo`


1. Изменить спецификацию **protos/echo.proto**
2. Изменить код обработчика в **sources/main.cpp**
3. Проверить работособность с помощью **curl**
 

## Links

[Protobuf](https://developers.google.com/protocol-buffers/docs/overview)
[gRPC gateway](https://github.com/grpc-ecosystem/grpc-gateway)
