package repl

import (
	"fmt"
)

func commandMap(cfg *Config, _ ...string) error {
	locationsPage, err := cfg.PokeClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsPage.Next
	cfg.previousLocationsURL = locationsPage.Previous

	for _, area := range locationsPage.Results {
		fmt.Printf("%s\n", area.Name)
	}

	return nil
}

func commandMapb(cfg *Config, _ ...string) error {
	if cfg.previousLocationsURL == nil {
		fmt.Println("You're on the first page")
		return nil
	}

	locationsPage, err := cfg.PokeClient.ListLocations(cfg.previousLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsPage.Next
	cfg.previousLocationsURL = locationsPage.Previous

	for _, area := range locationsPage.Results {
		fmt.Printf("%s\n", area.Name)
	}

	return nil
}
