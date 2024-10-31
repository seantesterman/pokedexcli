package main

import (
	"time"

	"github.com/seantesterman/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(time.Second * 1000),
	}

	startRepl(cfg)

}
