default: test

build:
	go build -v -o ./_bin/snake-game

run: build
	./_bin/snake-game

test:
	go test -v ./...
