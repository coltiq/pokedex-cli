package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/coltiq/pokedex-cli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	for {
		promptCommand := getPromptCommand("Pokedex >")
		if command, ok := getCliCommands()[promptCommand[0]]; ok {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
		fmt.Println("Unknown Command")
	}
}

func getPromptCommand(prompt string) []string {
	var commands []string
	r := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(prompt + " ")
		r.Scan()

		commands = cleanInput(r.Text())
		if len(commands[0]) != 0 {
			break
		}
	}
	return commands
}

func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCliCommands() map[string]cliCommand {
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
			description: "Display the names of the next 20 locations",
			callback:    commandMapF,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the names of the previous 20 locations",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Display the name of pokemon in given area",
			callback:    commandExplore,
		},
	}
}
