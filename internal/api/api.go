package api

import (
	"encoding/json"
	"fmt"
	"github.com/bevane/pokedex/internal/pokecache"
	"io"
	"net/http"
	"time"
)

var cache = pokecache.NewCache(5 * time.Second)

type LocationAreas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocations(url string) (LocationAreas, error) {
	locationAreas := LocationAreas{}
	var body []byte
	cachedResponse, ok := cache.Get(url)
	if ok {
		body = cachedResponse
	} else {
		res, err := http.Get(url)
		if err != nil {
			return locationAreas, err
		}
		body, err = io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			return locationAreas, err
		}
		if res.StatusCode != 200 {
			return locationAreas, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		cache.Add(url, body)
	}
	err := json.Unmarshal(body, &locationAreas)
	if err != nil {
		return locationAreas, err
	}
	return locationAreas, nil
}
