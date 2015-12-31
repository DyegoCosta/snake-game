package main

import "flag"

var cutefood bool

func main() {

	flag.BoolVar(&cutefood, "cutefood", true, "display a cute character as food")
	flag.Parse()

	g := NewGame()
	g.Start()
}
