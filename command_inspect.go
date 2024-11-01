package main

import (
	"errors"
	"fmt"
	"strings"
)

func callbackInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("No pokémon nameprovided")
	}
	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return errors.New("You haven't caught this pokémon yet")
	}
	fmt.Println("")
	fmt.Printf("Name: %s\n", strings.Title(pokemon.Name))
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("\nTypes:")
	for _, typ := range pokemon.Types {
		fmt.Printf(" - %s\n", typ.Type.Name)
	}
	fmt.Println("")
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("")

	return nil
}
