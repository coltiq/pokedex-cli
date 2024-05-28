package main

import (
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
    name := args[0]
    
    if pokemon, ok := cfg.pokedex[name]; ok {
        fmt.Printf("Name: %s\n", pokemon.Name)
        fmt.Printf("Height: %d\n", pokemon.Height)
        fmt.Printf("Weight: %d\n", pokemon.Weight)

        fmt.Println("Stats:")
        fmt.Printf("  -hp: %d\n", pokemon.Stats[0].BaseStat)
        fmt.Printf("  -attack: %d\n", pokemon.Stats[1].BaseStat)
        fmt.Printf("  -defense: %d\n", pokemon.Stats[2].BaseStat)
        fmt.Printf("  -special-attack: %d\n", pokemon.Stats[3].BaseStat)
        fmt.Printf("  -special-defense: %d\n", pokemon.Stats[4].BaseStat)
        fmt.Printf("  -speed: %d\n", pokemon.Stats[5].BaseStat)

        fmt.Println("Types: ")

        for _, ele := range pokemon.Types {
            fmt.Printf("  - %s\n", ele.Type.Name)
        }

        return nil
    }
    return fmt.Errorf("%s is not in Pokedex", name)
}
