package main

import (
    "fmt"
    api "github.com/coltiq/pokedex-cli/pokeapi"
)


func commandMap() error {
    fmt.Println(api.CommandGetLocations("https://pokeapi.co/api/v2/location/"))
    return nil
}


