package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {

	if len(args) != 1 {
		return fmt.Errorf("you must provide a pokemon name")
	}
	pokeName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokeName)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	random := rand.Intn(100)
	if random > (40 + (pokemon.BaseExperience / 10)) {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		cfg.pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}
