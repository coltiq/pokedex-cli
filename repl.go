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
		promptCommands := getPromptCommands("Pokedex >")

        commandName := promptCommands[0]
        args := []string{}
        if len(promptCommands) > 1 {
            args = promptCommands[1:]
        }

		if command, ok := getCliCommands()[commandName]; ok {
			if !checkUsage(command, args) {
				continue
			}
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
		fmt.Println("Unknown Command")
	}
}

func getPromptCommands(prompt string) []string {
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

func checkUsage(command cliCommand, extraCommands []string) bool {
	if len(extraCommands) != command.extraCommands {
		fmt.Println("Incorrect Usage:")
		fmt.Printf(" > %s\n", command.usage)
		return false
	}
	return true
}

type cliCommand struct {
	name          string
	description   string
	extraCommands int
	usage         string
	callback      func(*config, ...string) error
}

func getCliCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:          "help",
			description:   "Displays a help message",
			extraCommands: 0,
			usage:         "help [no extra commands]",
			callback:      commandHelp,
		},
		"exit": {
			name:          "exit",
			description:   "Exit the Pokedex",
			extraCommands: 0,
			usage:         "exit [no extra commands]",
			callback:      commandExit,
		},
		"map": {
			name:          "map",
			description:   "Display the names of the next 20 locations",
			extraCommands: 0,
			usage:         "map [no extra commands]",
			callback:      commandMapF,
		},
		"mapb": {
			name:          "mapb",
			description:   "Display the names of the previous 20 locations",
			extraCommands: 0,
			usage:         "mapb [no extra commands]",
			callback:      commandMapB,
		},
		"explore": {
			name:          "explore",
			description:   "Display the name of pokemon in given area",
			extraCommands: 1,
			usage:         "explore [location-area]",
			callback:      commandExplore,
		},
	}
}
