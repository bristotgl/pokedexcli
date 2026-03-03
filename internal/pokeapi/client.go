package pokeapi

import (
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
}

func NewClient(timout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timout,
		},
	}
}