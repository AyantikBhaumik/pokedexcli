package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(c *config) {
	reader := bufio.NewScanner(os.Stdin)

	callbackHelp(c, "")

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		words := cleanInput(reader.Text())
		if len(words)==0 {
			continue
		}
		commandName := words[0]
		command, exists := getCommands()[commandName]
		var param string
		if commandName=="explore" || commandName=="catch" || commandName=="inspect" {
			if len(words)!=2 {
				if commandName=="explore" {
					fmt.Println("Enter area name!!")
					fmt.Println()
				} else {
					fmt.Println("Enter pokemon name!!")
					fmt.Println()
				}
				continue
			} else {
				param = words[1]
			}
		}
		if exists {
			err := command.callback(c, param)
			if err!=nil {
				fmt.Println(err)
				fmt.Println()
			}
			continue
		} else {
			fmt.Println("Unknown command")
			fmt.Println()
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    callbackExit,
		},
		"map": {
			name:        "map",
			description: "Lists the next page of location areas",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists the previous page of location areas",
			callback:    callbackMapb,
		},
		"explore": {
			name:        "explore <area_name>",
			description: "Lists the pokemons in a particular area",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Tries to catch pokemon",
			callback:    callbackCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "Shows information about pokemon that has been caught",
			callback:    callbackInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Shows all the pokemons that has been caught",
			callback:    callbackPokedex,
		},
	}
}