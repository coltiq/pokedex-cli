package main

import (
	"os"
)

func commandExit(cfg *config, extraCommands []string) error {
	os.Exit(0)
	return nil
}
