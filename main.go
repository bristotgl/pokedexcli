package main

import (
	"time"

	"github.com/bristotgl/pokedexcli/internal/pokeapi"
	"github.com/bristotgl/pokedexcli/internal/pokecache"
	"github.com/bristotgl/pokedexcli/internal/repl"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cache := pokecache.NewCache(10 * time.Second)

	cfg := &repl.Config{
		PokeClient: pokeClient,
		PokeCache:  cache,
	}
	
	repl.StartRepl(cfg)
}
