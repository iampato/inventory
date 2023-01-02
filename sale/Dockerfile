FROM golang:1.19 AS build

# ENV CGO_ENABLED=0
RUN mkdir -p /go/src/inventory

ADD . /go/src/inventory

WORKDIR /go/src/inventory/sale

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run /go/src/inventory/sale/cmd  -watch inventory/sale
