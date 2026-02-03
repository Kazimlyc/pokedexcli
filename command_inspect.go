package main

import "fmt"

func commandInspect(cfg *config, args ...string) error {

	if len(args) != 1 {
		return fmt.Errorf("you must provide a pokemon name")
	}

	pokeName := args[0]

	pokemon, ok := cfg.pokedex[pokeName]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, s := range pokemon.Stats {
		fmt.Printf(" -%s: %d\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("Types:")
	for _, p := range pokemon.Types {
		fmt.Printf(" - %s\n", p.Type.Name)
	}
	return nil

}
