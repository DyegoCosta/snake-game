FROM golang:alpine as builder
ENV LANG en_US.UTF-8

WORKDIR /go/src/github.com/dyegocosta/snake-game
COPY . .

RUN apk add --no-cache git
RUN go get ./
RUN go build -ldflags="-extldflags=-static" -o /go/bin/snake-game

FROM scratch
WORKDIR /
COPY --from=builder /go/bin/snake-game .

ENTRYPOINT ["./snake-game"]
