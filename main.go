package main

import (
	"time"

	"github.com/Rifkiilmi/pokedexcli/internal/pokeapi"
)

type Config struct {
	pokeApiClient        pokeapi.Client
	nextLocationAreasURL *string
	prevLocationAreasURL *string
	caughtPokemon        map[string]pokeapi.Pokemon
}

func main() {
	cfg := Config{
		pokeApiClient: pokeapi.NewClient(time.Hour),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}

	startRepl(&cfg)
}
