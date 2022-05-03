FROM golang:1.18.1-alpine

ADD ./ /usr/bin/nitra-bot

WORKDIR /usr/bin/nitra-bot

RUN go mod download

ENTRYPOINT /usr/bin/nitra-bot

RUN go run ./cmd/main.go


