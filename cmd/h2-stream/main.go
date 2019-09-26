package main

import (
	"log"

	"github.com/igolaizola/h2-stream/internal/cli"
)

func main() {
	cmd := cli.NewCommand()

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
