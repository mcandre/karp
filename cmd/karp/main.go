package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mcandre/karp"
	"github.com/skratchdot/open-golang/open"
)

func usage() {
	fmt.Println("Usage: ", os.Args[0], "<path>")
	fmt.Println("-v Show version information")

	os.Exit(0)
}

func versionBanner() {
	fmt.Println(os.Args[0], karp.Version)

	os.Exit(0)
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}

	if os.Args[1] == "-v" || os.Args[1] == "--version" {
		versionBanner()
	}

	err := open.Run(os.Args[1])

	if err != nil {
		log.Panic(err)
	}
}
