FROM golang:1.19 AS build

ENV CGO_ENABLED=0
RUN mkdir -p /go/src/inventory

WORKDIR /go/src/inventory/user

ADD . /go/src/inventory
ADD config.ini /go/src/inventory/user/config.ini

# ADD config.ini user/config.ini

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run /go/src/inventory/user/cmd  -watch inventory/user
