package main

import "fmt"

func main() {
	for {
		inputCommand := stringPrompt("Pokedex >")
		commandMap := configure()

		if inputCommand == "help" {
			if cmd, exist := commandMap["help"]; exist {
				err := cmd.callback()
				if err != nil {
					fmt.Printf("Error executing 'help' command:%v\n", err)
				}
			}

		}

		if inputCommand == "exit" {
			if cmd, exist := commandMap["exit"]; exist {
				err := cmd.callback()
				if err != nil {
					fmt.Printf("Error executing 'exit' command:%v\n", err)
				}
			}
		}
	}
}
