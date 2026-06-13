package main

import (
	"fmt"
	"sort"
)

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")

	names := cfg.caughtPokemonNames
	if len(names) == 0 {
		names = make([]string, 0, len(cfg.pokedex))
		for name := range cfg.pokedex {
			names = append(names, name)
		}
		sort.Strings(names)
	}

	for _, name := range names {
		if _, ok := cfg.pokedex[name]; ok {
			fmt.Printf(" - %s\n", name)
		}
	}

	return nil
}
