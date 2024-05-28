package main

import (
    "errors"
    "fmt"
)

func commandExplore(cfg *config, args ...string) error {
	locationArea := &args[0]
	pokemonResp, err := cfg.pokeapiClient.GetPokemon(locationArea)
	if err != nil {
		return errors.New("Location Does Not Exist")
	}

	pokemonEncounters := pokemonResp.PokemonEncounters

	fmt.Printf("Exploring %s...\n", *locationArea)
	fmt.Println("Found Pokemon:")

	for _, enc := range pokemonEncounters {
		name := enc.Pokemon.Name
		fmt.Printf(" - %s\n", name)
	}

	return nil
}
