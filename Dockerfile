# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.21 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /server main.go
	# GOARCH=amd64 GOOS=linux go build -o ${BUILD_FOLDER}/latest/${BUILD_BINARY}-linux-amd64 main.go
# # Run the tests in the container
# FROM build-stage AS run-test-stage
# RUN go test -v ./...

# Deploy the application binary into a lean image
# FROM gcr.io/distroless/base-debian11 AS build-release-stage
FROM alpine:3.19 AS build-release-stage

WORKDIR /app

COPY ./data/*.txt /app/data/
COPY --from=build-stage /server /app/server

EXPOSE 3000

# USER nonroot:nonroot

ENTRYPOINT ["/app/server"]