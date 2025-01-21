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

func getLocations(url string) ([]string, error) {
	locationNames := []string{}
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	locationAreas := LocationAreas{}
	err = json.Unmarshal(body, locationAreas)
	if err != nil {
		return nil, err
	}
	for _, result := range locationAreas.Results {
		locationNames = append(locationNames, result.Name)
	}
	return locationNames, nil
}
