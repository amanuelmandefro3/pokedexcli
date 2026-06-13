package main

import (
	"github.com/amanuelmandefro3/pokedexcli/internal/pokeapi"
	"time"
)

func main() {
	pokeapiClient := pokeapi.NewClient(10*time.Second, 5*time.Second)
	cfg := &config{
		pokeapiClient: pokeapiClient,
		pokedex:       map[string]pokeapi.Pokemon{},
	}
	startRepl(cfg)
}
