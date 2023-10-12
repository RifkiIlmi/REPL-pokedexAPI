package main

import (
	"fmt"
)

func callbackPokedex(cfg *Config, args ...string) error {
	if len(cfg.caughtPokemon) < 1 {
		fmt.Println("you does not have pokemon yet, pelase catch!")
	}
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf("- %v\n", pokemon.Name)
	}

	return nil
}
