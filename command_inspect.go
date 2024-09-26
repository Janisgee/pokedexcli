package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a valid pokemon name")
	}
	inputPokemon := args[0]

	inspectedPokemon, exists := cfg.caughtPokemon[inputPokemon]
	if !exists {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", inspectedPokemon.Name)
	fmt.Printf("Height: %d\n", inspectedPokemon.Height)
	fmt.Printf("Weight: %d\n", inspectedPokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, val := range inspectedPokemon.Stats {
		fmt.Printf("   -%s: %d\n", val.Stat.Name, val.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, val := range inspectedPokemon.Types {
		fmt.Printf("   - %s\n", val.Type.Name)
	}

	return nil
}
