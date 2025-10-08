# BASE IMAGE
FROM golang:1.24.8-alpine3.21 AS build-base

WORKDIR /code

FROM alpine:3.21 AS base

WORKDIR /usr/local/bin

# BUILD PROJECT
FROM build-base AS go-deps

COPY go.mod go.sum ./
RUN go mod download

FROM go-deps AS build

ARG ARCH=amd64
ARG OS=linux

ENV GOOS=$OS
ENV GOARCH=$ARCH
ENV CGO_ENABLED=0

COPY . .

RUN go build -ldflags="-s -w" -o ./.build/api-subscription-service ./main.go

FROM base AS prod

WORKDIR /app

COPY --from=build /code/.build/api-subscription-service /usr/local/bin/api-subscription-service

EXPOSE 3000
EXPOSE 9090

ENTRYPOINT ["api-subscription-service"]

CMD ["run"]

