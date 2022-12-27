FROM golang

RUN mkdir -p /go/src/inventory

ADD . /go/src/inventory

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run inventory/sale/cmd  -watch inventory/sale
