package main

import "fmt"

func commandExplore(cfg *config, extraCommands []string) error {
    locationArea := &extraCommands[0]
    pokemonResp, err := cfg.pokeapiClient.GetPokemon(locationArea)
    if err != nil {
        return err 
    }

    pokemonEncounters := pokemonResp.PokemonEncounters

    fmt.Printf("Exploring %s...\n", *locationArea)
    fmt.Println("Found Pokemon:")

    for _, pokemon := range pokemonEncounters {
        name := pokemon.Pokemon.Name
        fmt.Printf(" - %s\n", name)
    }

    return nil
}
