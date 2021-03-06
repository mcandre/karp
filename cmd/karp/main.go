package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mcandre/karp"
)

// usage prints basic command line interface syntax for this application.
func usage() {
	fmt.Println("Usage: ", os.Args[0], "<path>")
	fmt.Println("-v Show version information")

	os.Exit(0)
}

// versionBanner prints a command line-accessible version number.
func versionBanner() {
	fmt.Println(os.Args[0], karp.Version)

	os.Exit(0)
}

// main is the entrypoint for this application.
func main() {
	if len(os.Args) != 2 {
		usage()
	}

	if os.Args[1] == "-v" || os.Args[1] == "--version" {
		versionBanner()
	}

	if err := karp.Launch(os.Args[1]); err != nil {
		log.Panic(err)
	}
}
