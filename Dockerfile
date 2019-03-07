FROM golang:latest

RUN mkdir -p /go/src/jumpcloud-hasher
WORKDIR /go/src/jumpcloud-hasher

ADD . /go/src/jumpcloud-hasher
WORKDIR /go/src/jumpcloud-hasher

RUN go get
EXPOSE 8080
