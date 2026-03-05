package pokeapi

import (
	"encoding/json"
	"net/http"

	"github.com/bristotgl/pokedexcli/internal/converter"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	if cachedPokemon, ok := c.cache.Get(url); ok {
		return converter.FromBytes[Pokemon](cachedPokemon)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	var pokemon Pokemon
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&pokemon); err != nil {
		return Pokemon{}, err
	}

	pokemonBytes, err := converter.ToBytes(pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, pokemonBytes)
	return pokemon, nil
}
