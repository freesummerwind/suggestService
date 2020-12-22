#include <algorithm>
#include <fstream>
#include <iostream>
#include <memory>
#include <mutex>
#include <string>
#include <shared_mutex>
#include <thread>

#include <grpcpp/grpcpp.h>
#include <grpcpp/health_check_service_interface.h>
#include <grpcpp/ext/proto_server_reflection_plugin.h>

#include <nlohmann/json.hpp>

#include "suggestService.grpc.pb.h"

using grpc::Server;
using grpc::ServerBuilder;
using grpc::ServerContext;
using grpc::Status;
using suggestService::SuggestRequest;
using SuggestService::Suggest;
using suggestService::SuggestResponse;
using suggestService::suggestService;

struct Suggestion {
    std::string id;
    std::string name;
    int cost = 0;
};

void from_json(const nlohmann::json& j, Suggest& s) {
    s.set_text(j.at("name").get<std::string>());
    s.set_position(j.at("cost").get<uint>());
}

class suggestServiceImpl final : public suggestService::Service {
public:
    void upgradeVariants() {
        std::ifstream file;
        while(true) {
            file.open("../suggestions.json");
            std::shared_lock <std::shared_mutex> lock(mutex);
            file >> variants;
            lock.unlock();
            file.close();
            std::this_thread::sleep_for(std::chrono::minutes(15));
        }
    }

private:
  Status Query(ServerContext* context, const SuggestRequest* request,
               SuggestResponse* response) override {
    std::thread upgrade(&upgradeVariants, this);
    upgrade.detach();
    google::protobuf::RepeatedPtrField<Suggest> goodVariants;
    std::shared_lock<std::shared_mutex> lock(mutex);
    goodVariants.Reserve(variants.size());
    for (const auto& variant : variants) {
        if (variant.at("id").get<std::string>() == request->query()) {
            goodVariants.Add(variant.get<Suggest>());
        }
    }
    lock.unlock();
    std::sort(goodVariants.begin(), goodVariants.end(),
            [](const Suggest& lhs, const Suggest& rhs) -> bool {
        return lhs.position() < rhs.position();
    });
    for (size_t i = 0; i < goodVariants.size(); ++i) {
        goodVariants[i].set_position(i);
    }
      *response->mutable_suggest_answer() = goodVariants;
    return Status::OK;
  }

  nlohmann::json variants;
  std::shared_mutex mutex;
};

void RunServer() {
  std::string server_address("0.0.0.0:9090");
  suggestServiceImpl service;

  grpc::EnableDefaultHealthCheckService(true);
  grpc::reflection::InitProtoReflectionServerBuilderPlugin();
  ServerBuilder builder;
  // Listen on the given address without any authentication mechanism.
  builder.AddListeningPort(server_address, grpc::InsecureServerCredentials());
  // Register "service" as the instance through which we'll communicate with
  // clients. In this case it corresponds to an *synchronous* service.
  builder.RegisterService(&service);
  // Finally assemble the server.
  std::unique_ptr<Server> server(builder.BuildAndStart());
  std::cout << "Server listening on " << server_address << std::endl;

  // Wait for the server to shutdown. Note that some other thread must be
  // responsible for shutting down the server for this call to ever return.
  server->Wait();
}

int main(int argc, char** argv) {
  RunServer();

  return 0;
}
