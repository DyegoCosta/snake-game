package main

import "flag"

var cutefood *bool

func main() {
	cutefood = flag.Bool("cutefood", true, "display a cute character as food")
	flag.Parse()

	g := NewGame()
	g.Start()
}
