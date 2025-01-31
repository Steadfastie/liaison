FROM golang:latest AS build
WORKDIR /app

COPY go/go.mod go/go.sum ./
RUN go mod download

COPY go/ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /out

# Deploy the application binary into a lean image
FROM gcr.io/distroless/static-debian12 AS release

WORKDIR /

COPY --from=build /out /out

EXPOSE 8080

USER nonroot:nonroot
ENTRYPOINT ["/out"]