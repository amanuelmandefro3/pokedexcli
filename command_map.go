package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, args ...string) error {
	locationRes, err := cfg.pokeapiClient.GetLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}

	cfg.nextLocationAreaURL = locationRes.Next
	cfg.previousLocationAreaURL = locationRes.Previous
	for _, loc := range locationRes.Results {
		fmt.Printf("%s\n", loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.previousLocationAreaURL == "" {
		return errors.New("you're on the first page")
	}

	locationRes, err := cfg.pokeapiClient.GetLocationAreas(cfg.previousLocationAreaURL)
	if err != nil {
		return err
	}

	cfg.nextLocationAreaURL = locationRes.Next
	cfg.previousLocationAreaURL = locationRes.Previous
	for _, loc := range locationRes.Results {
		fmt.Printf("%s\n", loc.Name)
	}
	return nil
}
