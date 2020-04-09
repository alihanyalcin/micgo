package main

import (
	"fmt"
	"github.com/alihanyalcin/micgo/generator"
	"os"
	"time"
)

func main() {

	start := time.Now()

	// Get args
	args := os.Args[1:]
	generator.Run(args)

	fmt.Println(time.Since(start))
}
