package main

import (
	"test/args"
	"test/repeat"
)

func main() {
	runArgsEcho()
	repeatSearch()
}

func runArgsEcho() {
	args.Echo()
}

func repeatSearch() {
	repeat.SearchV3()
}
