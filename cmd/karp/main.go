package main

import (
	"fmt"
	"os"

	"github.com/skratchdot/open-golang/open"
)

const version = "0.0.1"

func usage() {
	fmt.Printf("Usage: %s <path>\n\n", os.Args[0])
	fmt.Println("-v Show version information")

	os.Exit(0)
}

func versionBanner() {
	fmt.Printf("%s %s\n", os.Args[0], version)

	os.Exit(0)
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}

	if os.Args[1] == "-v" || os.Args[1] == "--version" {
		versionBanner()
	}

	open.Run(os.Args[1])
}
