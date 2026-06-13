package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("you must provide a location area")
	}
	locationRes, err := cfg.pokeapiClient.GetLocationArea(args[0])
	if err != nil {
		return err
	}

	for _, pokemon := range locationRes.PokemonEncounters {
		fmt.Printf("%s\n", pokemon.Pokemon.Name)
	}
	return nil
}
