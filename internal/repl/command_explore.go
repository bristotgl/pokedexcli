package repl

import (
	"fmt"
)

func commandExplore(cfg *Config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("You must provide a location name")
		return nil
	}

	locationName := args[0]

	fmt.Printf("Exploring %s...\n", locationName)
	locationArea, err := cfg.PokeClient.GetLocation(locationName)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
