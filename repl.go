package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Janisgee/pokedexcli/internal/pokeapi"
)

type Config struct {
	pokeapiClient        pokeapi.Client
	nextLocationsURL     *string
	previousLocationsURL *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, ...string) error
}

func startRepl(cfg *Config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		userInput := reader.Text()
		words := cleanInput(userInput)
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("unknown command")
			continue
		}

	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func getCommands() map[string]cliCommand {

	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display the name of next 20 locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the name of previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Find all Pokemon in a given area",
			callback:    commandExplore,
		},
	}

}
