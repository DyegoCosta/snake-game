FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download \
    && CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /snake-game .

FROM scratch
ENV LANG=en_US.UTF-8
COPY --from=builder /snake-game /snake-game
ENTRYPOINT ["/snake-game"]
