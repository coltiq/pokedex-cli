package main

import (
    "errors"
    "fmt"
    api "github.com/coltiq/pokedex-cli/pokeapi"
)


func commandMap() error {
    locations := api.CommandGetLocations(conf.Next)
    results := locations.Results
    for _, location := range results {
        fmt.Println(location.Name)
    }
    next := locations.Next
    previous := locations.Previous
    conf.Next = next
    conf.Previous = previous
    return nil
}


func commandMapB() error {
    if conf.Previous == "" {
        return errors.New("On first page")
    }
    locations := api.CommandGetLocations(conf.Previous)
    results := locations.Results
    for _, location := range results {
        fmt.Println(location.Name)
    }
    next := locations.Next
    previous := locations.Previous
    conf.Next = next
    conf.Previous = previous
    return nil
}

