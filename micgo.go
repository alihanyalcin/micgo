package main

import (
	"github.com/alihanyalcin/micgo/generator"
	"os"
)

func main() {
	// Get args
	args := os.Args[1:]
	generator.Run(args)
}
