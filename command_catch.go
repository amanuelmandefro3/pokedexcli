package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("you must provide a pokemon name")
	}
	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	pokemonRes, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	caught := pokemonRes.BaseExperience <= 0 || rand.Intn(pokemonRes.BaseExperience) < 50
	if caught {
		if _, ok := cfg.pokedex[pokemonRes.Name]; !ok {
			cfg.caughtPokemonNames = append(cfg.caughtPokemonNames, pokemonRes.Name)
		}
		cfg.pokedex[pokemonRes.Name] = pokemonRes
		fmt.Printf("%s was caught!\n", pokemonRes.Name)
	} else {
		fmt.Printf("%s escaped!\n", pokemonRes.Name)
	}
	return nil
}
