package main

import (
	"os"
)

func callbackExit(cfg *Config, args ...string) error {
	os.Exit(0)
	return nil

}
