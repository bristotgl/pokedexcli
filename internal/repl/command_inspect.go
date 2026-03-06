package repl

import "fmt"

func commandInspect(cfg *Config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("You must provide a pokemon name")
		return nil
	}

	pokemonName := args[0]
	pokemon, ok := cfg.CaughtPokemon[pokemonName]
	if !ok {
		fmt.Println("You haven't caught this pokemon yet!")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, pokemonType := range pokemon.Types {
		fmt.Printf(" -%s\n", pokemonType.Type.Name)
	}

	return nil
}
