FROM --platform=$BUILDPLATFORM mcr.microsoft.com/dotnet/sdk:9.0-alpine AS build
ARG TARGETARCH

RUN curl -Lo /usr/local/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.4.37/grpc_health_probe-linux-amd64 && \
    chmod +x /usr/local/bin/grpc_health_probe

RUN apk add --no-cache grpc-plugins
ENV PROTOBUF_PROTOC=/usr/bin/protoc \
    GRPC_PROTOC_PLUGIN=/usr/bin/grpc_csharp_plugin

COPY dotnet/service/*.csproj ./dotnet/service/
COPY dotnet/application/*.csproj ./dotnet/application/
COPY dotnet/infrastructure/*.csproj ./dotnet/infrastructure/
COPY proto/ ./proto/

RUN dotnet restore ./dotnet/service -a $TARGETARCH

COPY dotnet/service ./dotnet/service/
COPY dotnet/application ./dotnet/application/
COPY dotnet/infrastructure ./dotnet/infrastructure/

WORKDIR /dotnet/service

RUN dotnet publish --no-restore -c Release -a $TARGETARCH -o out

FROM mcr.microsoft.com/dotnet/aspnet:9.0-alpine AS runtime
EXPOSE 5002
WORKDIR /app
COPY --from=build /dotnet/service/out .
ENTRYPOINT ["dotnet", "service.dll"]