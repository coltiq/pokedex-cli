package main

import (
	"errors"
	"fmt"
)

func commandMapF(cfg *config, args ...string) error {
	locationsResp, err := cfg.pokeapiClient.GetLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	fmt.Println("Location Areas:")
	for _, location := range locationsResp.Results {
		fmt.Printf(" - %v\n", location.Name)
	}
	return nil
}

func commandMapB(cfg *config, args ...string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("You're on the first page")
	}

	locationsResp, err := cfg.pokeapiClient.GetLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	fmt.Println("Location Areas:")
	for _, location := range locationsResp.Results {
		fmt.Printf(" - %v\n", location.Name)
	}
	return nil
}
