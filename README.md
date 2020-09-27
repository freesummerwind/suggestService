# Лабораторная работа №7

## Задание

```sh
# install protobuf with gRPC C++ plugin
brew install protobuf
where protoc
git clone https://github.com/grpc/grpc --recursive
cd grpc
mkdir -p cmake/build
cd cmake/build
cmake ../..
make
mkdir $HOME/bin
cp grpc_cpp_plugin $HOME/bin
export PATH=$PATH:$HOME/bin
```
```sh
# install gRPC Go plugins
brew install golang
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

mkdir -p pkg/echo
```

```sh
# generate messages into pkg/echo.pb.go
protoc -I protos \
  --go_out ./pkg/echo \
  --go_opt plugins=grpc \
  --go_opt paths=source_relative \
  protos/echo.proto

# generate client/service into pkg/echo.pb.gw.go
protoc -I protos --grpc-gateway_out ./pkg/echo \
     --grpc-gateway_opt=logtostderr=true \
     --grpc-gateway_opt=paths=source_relative \
     --grpc-gateway_opt=generate_unbound_methods=true \
  protos/echo.proto
```

```sh
# build and run reverse proxy server: JSON <-> gRPC
go build main.go
./main
Server listening on 0.0.0.0:8080
```

```sh
# build and run gRPC server
./tools/polly/bin/polly
_builds/default/echoservice
Server listening on 0.0.0.0:9090
```

```sh
# send request to gRPC server via reverse proxy
curl --header "Content-Type: application/json" \
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
