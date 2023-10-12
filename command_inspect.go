package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("invalid pokemon\n")
	}
	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		fmt.Printf("you have'nt cought this pokemon yet\n")
		return nil
	}

	fmt.Printf("Name %s\n", pokemon.Name)
	fmt.Printf("Height %d\n", pokemon.Height)
	fmt.Printf("Weight %d\n", pokemon.Weight)
	fmt.Println("Stast:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Abilities:")
	for _, ability := range pokemon.Abilities {
		fmt.Printf(" - %v\n", ability.Ability.Name)
	}

	return nil
}
