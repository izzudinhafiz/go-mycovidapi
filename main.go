package main

import (
	"fmt"
	"os"

	"github.com/izzudinhafiz/go-mycovidapi/ingest"
	"github.com/izzudinhafiz/go-mycovidapi/server"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Printf("Requires argument to run program. either `serve` or `ingest`")
		return
	}

	if args[0] == "ingest" {
		ingest.Ingest()
		return
	} else if args[0] != "serve" {
		fmt.Printf("Unknown arguments to program. Can be either `serve` or `ingest`. Got %v\n", args)
		return
	}

	s := server.New()
	s.Run()
}