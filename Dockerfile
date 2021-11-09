# syntax=docker/dockerfile:1
# source: https://docs.docker.com/language/golang/build-images/
# Build the go module

FROM golang:1.17-buster AS build

LABEL author="idebrody"

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o /go-web-api

# Deploy the go module to a slim image
FROM gcr.io/distroless/base-debian10:latest

#
ENV GOROUTINETRESHOLD=100

WORKDIR /

COPY --from=build /go-web-api /go-web-api

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/go-web-api"]