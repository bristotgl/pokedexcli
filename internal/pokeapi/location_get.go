package pokeapi

import (
	"encoding/json"
	"net/http"

	"github.com/bristotgl/pokedexcli/internal/converter"
)

func (c *Client) GetLocation(locationName string) (LocationArea, error) {
	url := baseURL + "/location-area/" + locationName

	if val, ok := c.cache.Get(url); ok {
		return converter.FromBytes[LocationArea](val)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	var locationArea LocationArea
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&locationArea); err != nil {
		return LocationArea{}, err
	}

	encodedLocation, err := converter.ToBytes(locationArea)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(url, encodedLocation)
	return locationArea, nil
}
