package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a valid pokemon name")
	}
	inputName := args[0]

	pokemonResp, err := cfg.pokeapiClient.GetPokemon(inputName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonResp.Name)

	//Calculate chance to catch a pokemon
	maxBaseExp := 635
	baseExp := pokemonResp.BaseExperience
	catchChance := 100 - ((baseExp * 100) / maxBaseExp)
	randomNumber := rand.IntN(100)

	//Determine if user can catch the pokemon
	//Sucess
	if randomNumber < catchChance {
		fmt.Printf("%s was caught!\n", pokemonResp.Name)
		cfg.caughtPokemon[pokemonResp.Name] = pokemonResp
	} else {
		//Fail
		fmt.Printf("%s escaped!\n", pokemonResp.Name)
	}

	return nil
}
