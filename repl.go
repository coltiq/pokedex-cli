package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func startRepl() {
    cliCommands := getCliCommands()
    for {
        promptCommand := getPromptCommand("Pokedex >")
        if value, ok := cliCommands[promptCommand]; ok{
            err := value.callback()
            if err != nil || promptCommand == "exit" {
                break
            }
            continue
        }
        fmt.Println("Command Does Not Exist")
    }
}


func getPromptCommand(prompt string) string {
    var command string
    r := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print(prompt+" ")
        r.Scan()

        command = cleanInput(r.Text())
        if len(command) != 0 {
            break
        }
    }
    return command
}


func cleanInput(text string) string {
    command := strings.ToLower(text)
    return command
}


type cliCommand struct {
    name        string
    description string
    callback    func() error
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
    }
}


