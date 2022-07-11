# syntax=docker/dockerfile:1

########################
# Build
########################

FROM golang:1.17 AS build

WORKDIR /app

COPY . .
RUN go mod download

RUN go build -o /sobrian ./cmd/main.go

########################
# Deploy
########################
FROM debian:stable-slim

WORKDIR /

COPY --from=build /sobrian /sobrian

ENV port=8080

EXPOSE 8080

# USER nonroot:nonroot

ENTRYPOINT ["/sobrian"]