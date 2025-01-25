![Go](https://github.com/Steadfastie/liaison/actions/workflows/go.yml/badge.svg?branch=main) ![.NET](https://github.com/Steadfastie/liaison/actions/workflows/dotnet.yml/badge.svg)


![liason logo _ open](https://github.com/user-attachments/assets/a1b9b949-146a-48ca-b929-7114915517e2)
# 🦉 liaison

Go client generate command (WSL-compatible):

     protoc --proto_path=./proto --go_out=./go/client --go_opt=paths=source_relative \
     --go-grpc_out=./go/client --go-grpc_opt=paths=source_relative \
     --experimental_allow_proto3_optional \
     ./proto/order_service.proto ./proto/order_request.proto ./proto/order_response.proto

Docker build dotnet

     docker build -f dotnet.dockerfile ../
