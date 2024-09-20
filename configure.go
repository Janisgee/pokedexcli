package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandHelp() error {
	cmd := configure()
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Printf("%s: %s\n", cmd["help"].name, cmd["help"].description)
	fmt.Printf("%s: %s\n", cmd["exit"].name, cmd["exit"].description)
	fmt.Println()
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func stringPrompt(label string) string {
	var inputString string
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+" ")
		inputString, _ = reader.ReadString('\n')
		if inputString != "" {
			break
		}

	}
	inputString = strings.ToLower(inputString)
	return strings.TrimSpace(inputString)
}

func configure() map[string]cliCommand {

	return map[string]cliCommand{"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

}
