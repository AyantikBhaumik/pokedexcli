package main

import (
	"time"
	"github.com/AyantikBhaumik/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	pokedex map[string]pokeapi.Pokemon
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		pokedex: make(map[string]pokeapi.Pokemon),
	}
	startRepl(&cfg)
}