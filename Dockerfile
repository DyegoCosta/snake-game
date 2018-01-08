# Build Stage
FROM golang:alpine AS build-env
ADD . /src
RUN apk update && \
      apk add git && \
      cd /src && \
      go get github.com/DyegoCosta/snake-game/snake && \
      go build -v -o ./_bin/snake-game

# Final Stage
FROM alpine
ENV LANG en_US.UTF-8
WORKDIR /app
COPY --from=build-env /src/_bin/snake-game /app/
ENTRYPOINT ./snake-game
