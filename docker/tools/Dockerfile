FROM golang:1.17-alpine

ENV GO111MODULE=on

WORKDIR /app


RUN apk add --no-cache git make bash build-base protoc curl mc

#RUN git clone https://github.com/googleapis/googleapis.git /var/api/googleapis/google

RUN mkdir /var/proto && cd /var/proto && \
  curl -OL https://github.com/google/protobuf/releases/download/v3.2.0/protoc-3.2.0-linux-x86_64.zip && \
  unzip protoc-3.2.0-linux-x86_64.zip -d protoc3

RUN go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway && \
    go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger && \
    go get github.com/golang/protobuf/protoc-gen-go && \
    go get google.golang.org/grpc/cmd/protoc-gen-go-grpc

