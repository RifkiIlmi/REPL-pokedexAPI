package main

import (
	"errors"
	"fmt"
)

func callbackMap(cfg *Config, args ...string) error {
	resp, err := cfg.pokeApiClient.ListLocationAreas(cfg.nextLocationAreasURL)
	if err != nil {
		return err
	}

	fmt.Println("Location Areas:")
	for _, l := range resp.Results {
		fmt.Printf(" -. %s\n", l.Name)
	}

	cfg.nextLocationAreasURL = resp.Next
	cfg.prevLocationAreasURL = resp.Previous

	return nil
}

func callbackMapBack(cfg *Config, args ...string) error {
	if cfg.prevLocationAreasURL == nil {
		return errors.New("you are in the first page\n")
	}
	resp, err := cfg.pokeApiClient.ListLocationAreas(cfg.prevLocationAreasURL)
	if err != nil {
		return err
	}

	fmt.Println("Location Areas:")
	for _, l := range resp.Results {
		fmt.Printf(" -. %s\n", l.Name)
	}

	cfg.nextLocationAreasURL = resp.Next
	cfg.prevLocationAreasURL = resp.Previous

	return nil
}
