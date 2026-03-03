package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bristotgl/pokedexcli/internal/pokeapi"
	"github.com/bristotgl/pokedexcli/internal/pokecache"
)

type Config struct {
	nextLocationsURL     *string
	previousLocationsURL *string
	PokeClient           pokeapi.Client
	PokeCache            *pokecache.Cache
}

func StartRepl(cfg *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\nPokedex > ")
		scanner.Scan()

		if scanner.Err() != nil {
			os.Exit(1)
		}

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		command, exists := getCommands()[words[0]]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		if err := command.callback(cfg); err != nil {
			fmt.Println(err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *Config) error
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
			description: "Displays the names of the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 locations",
			callback:    commandMapb,
		},
	}
}
