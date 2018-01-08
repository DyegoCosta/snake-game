FROM golang:alpine
ENV LANG en_US.UTF-8

COPY . /go/src/github.com/dyegocosta/snake-game
WORKDIR /go/src/github.com/dyegocosta/snake-game

RUN apk add --no-cache git
RUN go get ./
RUN go build -o ./_bin/snake-game

ENTRYPOINT "./_bin/snake-game"
