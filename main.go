package main

import (
	"os"

	"github.com/lisabestteam/password-svc/internal/cli"
)

func main() {
	if !cli.Run(os.Args[:]) {
		os.Exit(1)
	}
}
