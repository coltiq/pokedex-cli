package main

import "fmt"

func commandExplore(cfg *config, extraCommands []string) error {
    locationArea := extraCommands[0]
    pokemonResp, err := cfg.pokeapiClient.GetPokemon(locationArea)
    if err != nil {
        return err 
    }

    fmt.Println(pokemonResp)

    return nil
}
