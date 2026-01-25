package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Next     *string
	Previous *string
}

func main() {
	cfg := &config{}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		input = strings.Fields(input)[0]
		input = strings.ToLower(input)
		command, exists := getCommands()[input]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}
		if err := command.callback(cfg); err != nil {
			fmt.Println(err)
		}
	}
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(cfg *config) error {
	type LocationAreasResp struct {
		Count    int
		Next     *string
		Previous *string
		Results  []struct {
			Name string
			URL  string
		}
	}
	// Todo access just name field and set cfg.next and previous
	data, err := FetchLocations("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		return err
	}
	var locationData LocationAreasResp
	if err := json.Unmarshal(data, &locationData); err != nil {
		return err
	}
	for i := 0; i < 20; i++ {
		fmt.Println(locationData.Results[i])
	}

	return nil
}

func commandMapb(cfg *config) error {

	return nil
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
	}
}
