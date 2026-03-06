package repl

import "fmt"

func commandPokedex(cfg *Config, _ ...string) error {
	if len(cfg.CaughtPokemon) == 0 {
		fmt.Println("Your pokedex is empty. Go catch'em all!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for name, _ := range cfg.CaughtPokemon {
		fmt.Printf("  -%s\n", name)
	}

	return nil
}
