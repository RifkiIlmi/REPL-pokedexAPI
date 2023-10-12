package main

import "fmt"

func callbackHelp(cfg *Config, args ...string) error {
	fmt.Println("Welcome to the Pokedex CLI help Menu")
	fmt.Println("Here are your available commands:")

	availableCommands := getCommands()
	for _, cmd := range availableCommands {
		fmt.Printf("- %s: %s\n", cmd.Name, cmd.Description)
	}

	return nil
}
