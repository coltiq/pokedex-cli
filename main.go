package main

import (
	"time"

	"github.com/coltiq/pokedex-cli/internal/pokeapi"
	//"github.com/coltiq/pokedex-cli/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
    baseURL := pokeapi.BaseURL
	//pokeCache := pokecache.NewCache(5 * time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
		//pokeCache: pokeCache,
        nextLocationsURL: &baseURL,
	}

	startRepl(cfg)
}
