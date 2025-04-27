package main

import (
	"os"

	"github.com/xcr-19/m365recon/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
