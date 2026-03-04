package main

import (
	"time"

	"github.com/bristotgl/pokedexcli/internal/pokeapi"
	"github.com/bristotgl/pokedexcli/internal/repl"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 10*time.Minute)
	cfg := &repl.Config{PokeClient: pokeClient}
	repl.StartRepl(cfg)
}
