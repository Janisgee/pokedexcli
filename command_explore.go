package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a valid location name")
	}
	inputLocation := args[0]

	locationResp, err := cfg.pokeapiClient.GetLocation(inputLocation)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", inputLocation)
	fmt.Println("Found Pokemon:")

	for _, encounters := range locationResp.PokemonEncounters {
		fmt.Printf(" - %s\n", encounters.Pokemon.Name)
	}

	return nil
}
