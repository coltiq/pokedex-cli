package main

import (
	"errors"
	"fmt"
)

func commandMapF(cfg *config) error {
    fmt.Printf("%v", *cfg.nextLocationsURL)
    if cfg.nextLocationsURL == nil {
        return errors.New("You are on the last page")
    }

	locationsResp, err := cfg.pokeapiClient.GetLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}
    fmt.Println(locationsResp)

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapB(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("You're on the first page")
	}

	locationsResp, err := cfg.pokeapiClient.GetLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}
	return nil
}
