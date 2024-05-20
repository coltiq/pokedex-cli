module github.com/coltiq/pokedex-cli

go 1.22.2

require packages/pokeapi/ v1.0.0
replace package/pokeapi/ => ./packages/pokeapi
