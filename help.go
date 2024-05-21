package main

import (
    "fmt"
)


func commandHelp(cfg *config) error{
    fmt.Println()
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:")
    fmt.Println()
    for _, cmd := range getCliCommands() {
        fmt.Printf("%v: %v\n", cmd.name, cmd.description)
    }
    fmt.Println()
    return nil
}
