package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *Config, args ...string) error {

	if len(cfg.caughtPokemon) == 0 {
		return errors.New("there is no pokemon have caught by user")
	}

	fmt.Println("Your Pokedex:")

	for pokemonName := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", pokemonName)
	}
	return nil
}
