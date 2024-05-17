package main

import (
    "bufio"
    "errors"
    "fmt"
    "os"
    "strings"
)

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


func getPromptCommand(prompt string) string {
    var command string
    r := bufio.NewReader(os.Stdin)
    for {
        fmt.Fprint(os.Stderr, prompt+" ")
        command, _ = r.ReadString('\n')
        if command != "" {
            break
        }
    }
    return strings.TrimSpace(command)
}

func commandHelp() error{
    cliCommands := getCliCommands()
    fmt.Println("\nWelcome to the Pokedex!\nUsage:\n")
    for _, command := range cliCommands {
        fmt.Printf("%v: %v\n", command.name, command.description)
    }
    fmt.Println("")
    return nil
}


func commandExit() error{
    return errors.New("Exiting Gracefully")
}


func main() {
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












