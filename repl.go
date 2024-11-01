package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex >")

		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Invalid Command")
			continue
		}
	}

}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{

		"catch": {
			name:        "catch {pokemon_name}",
			description: "Attempt to catch a pokémon and add it to your pokédex",
			callback:    callbackCatch,
		},
		"explore": {
			name:        "explore {location_area}",
			description: "Lists the pokémon in a loication area",
			callback:    callbackExplore,
		},
		"help": {
			name:        "help",
			description: "Prints the help menu",
			callback:    callbackHelp,
		},
		"inspect": {
			name:        "inspect {pokémon_name}",
			description: "Displays information about the chosen pokémon",
			callback:    callbackInspect,
		},
		"map": {
			name:        "map",
			description: "Lists some location areas",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Goes to previous map page",
			callback:    callbackMapb,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Views all the pokémon in your pokédex",
			callback:    callbackPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Turns off the Pokédex",
			callback:    callbackExit,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
