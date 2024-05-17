package main

import (
    "fmt"
)


func commandHelp() error{
    cliCommands := getCliCommands()
    fmt.Println("\nWelcome to the Pokedex!\nUsage:\n")
    for _, command := range cliCommands {
        fmt.Printf("%v: %v\n", command.name, command.description)
    }
    fmt.Println("")
    return nil
}
