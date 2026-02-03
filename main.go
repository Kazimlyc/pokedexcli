package main

import (
	"time"

	"github.com/Kazimlyc/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	pokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
	pokedex       map[string]pokeapi.Pokemon
}

func main() {
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(
			5*time.Second,
			5*time.Minute,
		),
		pokedex: make(map[string]pokeapi.Pokemon),
	}
	startRepl(cfg)

}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of 20 location areas in the Pokemon world",
			callback:    commandMapb,
		},
		"explore": {
			name: "explore",
			// Todo
			description: "",
			callback:    commandExplore,
		},
		"catch": {
			name: "catch",
			//todo
			description: "",
			callback:    commandCatch,
		},
	}
}
