package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bristotgl/pokedexcli/internal/pokeapi"
)

type Config struct {
	nextLocationsURL     *string
	previousLocationsURL *string
	PokeClient           pokeapi.Client
}

func StartRepl(cfg *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\nPokedex > ")
		scanner.Scan()
		fmt.Println()

		if scanner.Err() != nil {
			os.Exit(1)
		}

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		command, exists := getCommands()[commandName]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		args := words[1:]
		if err := command.callback(cfg, args...); err != nil {
			fmt.Println(err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *Config, args ...string) error
}

func cleanInput(text string) []string {
	lowercase := strings.ToLower(text)
	words := strings.Fields(lowercase)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore a location",
			callback:    commandExplore,
		},
	}
}
