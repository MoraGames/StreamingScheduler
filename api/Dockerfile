# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.18-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY internal/ ./internal/

RUN go mod download

COPY . .

RUN go build -o /api

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /api /api

EXPOSE 5001

#USER nonroot:nonroot

ENTRYPOINT ["/api"]
