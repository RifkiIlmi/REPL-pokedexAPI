package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon found\n")
	}
	pokemonName := args[0]

	pokemon, err := cfg.pokeApiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	const treshold = 60
	randNum := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("%v - %v - %v\n", pokemon.BaseExperience, randNum, treshold)

	if randNum >= treshold {
		fmt.Printf("failed to catch %s\n", pokemon.Name)
		return nil
	}

	cfg.caughtPokemon[pokemonName] = pokemon
	fmt.Printf("succes catch %s\n", pokemon.Name)

	return nil
}
