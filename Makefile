default: run

build:
	go build -v -o ./_bin/snake-game

run: build
	./_bin/snake-game
