package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must provide a location name")
	}
	areaName := args[0]

	location, err := cfg.pokeapiClient.GetLocation(areaName)
	if err != nil {
		return err
	}
	fmt.Printf("Explorin %s...\n", location.Name)
	fmt.Println("Found Pokemon:")

	for _, enc := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}

	return nil
}
