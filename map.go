package main

import (
    "pokedex-cli/packages/pokeapi"
)


func commandMap() error {
    fmt.Println(PokeAPI("get", "https://pokeapi.co/api/v2/location/"))
    return nil
}


