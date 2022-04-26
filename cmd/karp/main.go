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
	fmt.Println()
	fmt.Println("-h Show usage information")
	fmt.Println("-v Show version information")
}

// versionBanner prints a command line-accessible version number.
func versionBanner() {
	fmt.Println(os.Args[0], karp.Version)
}

// main is the entrypoint for this application.
func main() {
	if len(os.Args) == 0 {
		log.Fatalf("error: missing program name\n")
	}

	if len(os.Args) != 2 {
		usage()
		os.Exit(1)
	}

	if os.Args[1] == "-h" {
		usage()
		os.Exit(0)
	}

	if os.Args[1] == "-v" {
		versionBanner()
		os.Exit(0)
	}

	if err := karp.Launch(os.Args[1]); err != nil {
		os.Exit(1)
	}
}
