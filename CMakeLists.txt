cmake_minimum_required(VERSION 3.4)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

option(BUILD_COVERAGE "Build coverage" OFF)

set(
  HUNTER_CACHE_SERVERS
  "https://github.com/bmstu-iu8-cpp-sem-3/hunter-cache"
  CACHE STRING "Default cache server"
)

include("tools/gate/cmake/HunterGate.cmake")

HunterGate(
  URL "https://github.com/cpp-pm/hunter/archive/v0.23.218.tar.gz"
  SHA1 "9a3594a61227967fb92159339ba326701b287b19"
)

project(echoservice)
string(APPEND CMAKE_CXX_FLAGS " -pedantic -Werror -Wall -Wextra")
string(APPEND CMAKE_CXX_FLAGS " -Wno-unused-command-line-argument")
string(APPEND CMAKE_CXX_FLAGS " -Wshadow -Wnon-virtual-dtor")

hunter_add_package(GTest)
find_package(GTest CONFIG REQUIRED)

hunter_add_package(gRPC)
find_package(gRPC CONFIG REQUIRED)

hunter_add_package(Protobuf)
find_package(Protobuf CONFIG REQUIRED)

set(PROTO_GENERATED_SOURCES "${CMAKE_CURRENT_BINARY_DIR}/echo.pb.cc")
set(PROTO_GENERATED_HEADERS "${CMAKE_CURRENT_BINARY_DIR}/echo.pb.h")
set(PROTO_GENERATED_GRPC_SOURCES "${CMAKE_CURRENT_BINARY_DIR}/echo.grpc.pb.cc")
set(PROTO_GENERATED_GRPC_HEADERS "${CMAKE_CURRENT_BINARY_DIR}/echo.grpc.pb.h")

find_program(PROTOBUF_PROTOC protoc)
find_program(GRPC_CPP_PLUGIN_EXECUTABLE grpc_cpp_plugin)

get_filename_component(PROTO_FILE_PATH "./protos/echo.proto" ABSOLUTE)
get_filename_component(PROTO_INCLUDE_DIR "${PROTO_FILE_PATH}" PATH)
add_custom_command(
      OUTPUT  "${PROTO_GENERATED_SOURCES}"
              "${PROTO_GENERATED_HEADERS}"
              "${PROTO_GENERATED_GRPC_SOURCES}"
              "${PROTO_GENERATED_GRPC_HEADERS}"
      COMMAND protobuf::protoc
      ARGS --grpc_out "${CMAKE_CURRENT_BINARY_DIR}"
        --cpp_out "${CMAKE_CURRENT_BINARY_DIR}"
        -I"${PROTO_INCLUDE_DIR}"
        --plugin=protoc-gen-grpc="${GRPC_CPP_PLUGIN_EXECUTABLE}"
        "${PROTO_FILE_PATH}"
      DEPENDS "${PROTO_FILE_PATH}")

include_directories("${CMAKE_CURRENT_BINARY_DIR}")

add_executable(echoservice ${CMAKE_CURRENT_SOURCE_DIR}/sources/main.cpp
  ${PROTO_GENERATED_SOURCES}
  ${PROTO_GENERATED_GRPC_SOURCES})

target_link_libraries(echoservice
  gRPC::grpc
  protobuf::libprotobuf)
