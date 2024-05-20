package pokeapi

import "strings"


func PokeAPI(command string, URL string) {
    if strings.ToLower(command) == "get" {
        CommandGet(URL)
    }
}
