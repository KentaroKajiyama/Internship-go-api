FROM golang:1.21.5-bullseye as base

ENV TZ=Asia/Tokyo

RUN apt-get update &&\
    apt-get upgrade -y &&\
    apt-get install -y libicu-dev &&\
    apt-get clean &&\
    rm -rf /var/lib/apt/lists/*

WORKDIR /go/src

RUN go install github.com/pressly/goose/v3/cmd/goose@latest&&\
    go install github.com/google/wire/cmd/wire@latest &&\
    go install github.com/swaggo/swag/cmd/swag@latest &&\
    go install github.com/cespare/reflex@latest &&\
    go install github.com/go-delve/delve/cmd/dlv@latest &&\
    go install honnef.co/go/tools/cmd/staticcheck@latest

COPY ./go.mod /go/src
COPY ./go.sum /go/src

RUN go mod download

COPY ./ /go/src

RUN go build -o /go/bin/app /go/src/cmd/rest/main.go

CMD ["/go/bin/app"]

