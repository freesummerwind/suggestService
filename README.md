# Лабораторная работа №7

## Задание

```sh
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
