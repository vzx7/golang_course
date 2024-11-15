package main

import (
	"test/animation"
	"test/args"
	"test/repeat"
)

func main() {
	runArgsEcho()
	//repeatSearch()
	animation.Gen()
}

func runArgsEcho() {
	args.Echo()
}

func repeatSearch() {
	repeat.SearchV3()
}
