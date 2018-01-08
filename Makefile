default: test

build:
	go build -v -o ./_bin/snake-game

run: build
	./_bin/snake-game

run_on_docker:
	docker build -t snake-game . && docker run --rm -ti snake-game

test:
	go test -v ./...
