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

type LocationAreaDetails struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

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

func GetLocationDetails(locationName string) (LocationAreaDetails, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%v/", locationName)
	locationAreaDetails := LocationAreaDetails{}
	var body []byte
	cachedResponse, ok := cache.Get(url)
	if ok {
		body = cachedResponse
	} else {
		res, err := http.Get(url)
		if err != nil {
			return locationAreaDetails, err
		}
		body, err = io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			return locationAreaDetails, err
		}
		if res.StatusCode != 200 {
			return locationAreaDetails, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		cache.Add(url, body)
	}
	err := json.Unmarshal(body, &locationAreaDetails)
	if err != nil {
		return locationAreaDetails, err
	}
	return locationAreaDetails, nil
}
