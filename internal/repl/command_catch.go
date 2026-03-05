package repl

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *Config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("You must provide a pokemon name")
		return nil
	}

	pokemonName := args[0]
	pokemon, err := cfg.PokeClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	catchScore := rand.Intn(pokemon.BaseExperience)
	if catchScore > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)

	cfg.CaughtPokemon[pokemon.Name] = pokemon
	return nil
}
