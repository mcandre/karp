package main

import (
	"fmt"
	"os"

	"github.com/skratchdot/open-golang/open"
)

func usage() {
	fmt.Println("Usage: %s <file>", os.Args[0])
	os.Exit(0)
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}

	open.Run(os.Args[1])
}
