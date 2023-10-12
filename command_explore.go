package main

import (
	"errors"
	"fmt"
)

func callbackExplore(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area provided\n")
	}
	locationName := args[0]

	resp, err := cfg.pokeApiClient.GetLocationArea(locationName)
	if err != nil {
		return err
	}

	fmt.Printf("Pokemon in %s:\n", resp.Name)
	for _, v := range resp.PokemonEncounters {
		fmt.Printf(" -. %s\n", v.Pokemon.Name)
	}

	return nil
}
