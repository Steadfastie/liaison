![Go](https://github.com/Steadfastie/liaison/actions/workflows/go.yml/badge.svg?branch=main) ![.NET](https://github.com/Steadfastie/liaison/actions/workflows/dotnet.yml/badge.svg)

![liason logo _ open](https://github.com/user-attachments/assets/a1b9b949-146a-48ca-b929-7114915517e2)

# 🦉 liaison

This project was a part of an exploration of the gRPC framework. The main goal was to understand how gRPC works across different languages and to compare the development experience in each

It covered:

- Writing a service & testing it with the client in Go
- Writing a service & testing it with the client in .NET
- Trying out [Chiseled](https://devblogs.microsoft.com/dotnet/announcing-dotnet-chiseled-containers/) containers (.NET) and comparing them to [Distroless](https://github.com/GoogleContainerTools/distroless) (Go)

### 🗝️ Key learnings/findings
- gRPC is not scary
- Standardised code generation is 🚀 
- The size difference between essentially identical services in Go/Distroless & .NET/Chiseled is striking

![image](https://github.com/user-attachments/assets/5a3a0c0f-6e31-4e36-8086-765d391f1d98)

### 👩‍💻 Useful commands
Go proto generate command for Linux (WSL):
```
protoc --proto_path=./proto --go_out=./go/generated/client --go_opt=paths=source_relative \
--go-grpc_out=./go/generated/client --go-grpc_opt=paths=source_relative \
--experimental_allow_proto3_optional \
./proto/order*.proto
```                                       
```
protoc --proto_path=./proto --go_out=./go/generated/service --go_opt=paths=source_relative \
--go-grpc_out=./go/generated/service --go-grpc_opt=paths=source_relative \
--experimental_allow_proto3_optional \
./proto/tracking*.proto
```
Docker compose
```
docker compose -p liaison up -d
```

### 🍿 Handy links
- [About NotImplemented](https://github.com/grpc/grpc-go/issues/3794)
- [About versioning](https://learn.microsoft.com/en-us/aspnet/core/grpc/versioning?view=aspnetcore-9.0)
- [About health probes](https://github.com/grpc-ecosystem/grpc-health-probe)
