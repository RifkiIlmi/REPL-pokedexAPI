package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *Config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(" > ")

		scanner.Scan()
		text := scanner.Text()

		if len(text) == 0 {
			continue
		}

		cleaned := cleanInput(text)
		commandName := cleaned[0]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		availableCommands := getCommands()
		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Invalid Command")
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Print(err)
		}
	}
}

type cliCommands struct {
	Name        string
	Description string
	callback    func(*Config, ...string) error
}

func getCommands() map[string]cliCommands {
	return map[string]cliCommands{
		"help": {
			Name:        "help",
			Description: "Prints the help command",
			callback:    callbackHelp,
		},
		"map": {
			Name:        "map",
			Description: "List of Areas - Next page",
			callback:    callbackMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "List of Areas - Previous page",
			callback:    callbackMapBack,
		},
		"explore": {
			Name:        "explore { location_area }",
			Description: "List of Pokemons on the specific area",
			callback:    callbackExplore,
		},
		"catch": {
			Name:        "catch { pokemon_name }",
			Description: "Catch Pokemon",
			callback:    callbackCatch,
		},
		"inspect": {
			Name:        "inspect { pokemon_name }",
			Description: "Print detail about the pokemon",
			callback:    callbackInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Show all caught pokemon",
			callback:    callbackPokedex,
		},
		"exit": {
			Name:        "exit",
			Description: "Turn off the PokedexCLI",
			callback:    callbackExit,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
