package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
    name := &args[0]
    pokemon, err := cfg.pokeapiClient.GetPokemon(name)
    if err != nil {
        return errors.New("Pokemon Does Not Exist")
    }

    result := rand.Intn(pokemon.BaseExperience)

    fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
    if result > 40 {
            fmt.Printf("%s escaped\n", pokemon.Name)
            return nil
    }

    fmt.Printf("%s was caught!\n", pokemon.Name)

    cfg.pokedex[pokemon.Name] = pokemon
    return nil
}
