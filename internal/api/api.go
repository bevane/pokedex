package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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
	res, err := http.Get(url)
	if err != nil {
		return locationAreas, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return locationAreas, err
	}
	if res.StatusCode != 200 {
		return locationAreas, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	err = json.Unmarshal(body, &locationAreas)
	if err != nil {
		return locationAreas, err
	}
	return locationAreas, nil
}
