package main

import (
	"os"

	"github.com/catatsuy/lls/cli"
)

func main() {
	c := cli.NewCLI(os.Stdout, os.Stderr)
	os.Exit(c.Run(os.Args))
}
