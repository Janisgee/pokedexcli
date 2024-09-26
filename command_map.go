package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *Config, args ...string) error {
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.previousLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(cfg *Config, args ...string) error {
	if cfg.previousLocationsURL == nil {
		return errors.New("this is the first page")
	}
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.previousLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.previousLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
