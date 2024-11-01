package main

import (
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {
	fmt.Println("Pokémon in pokédex:")
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("You have not caught any pokémon yet!")
	}
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
