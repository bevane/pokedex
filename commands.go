package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/bevane/pokedex/internal/api"
)

func commandHelp(config *config, args ...string) error {
	fmt.Print("\nWelcome to the Pokedex!\nUsage:\n\n")
	commandMap := getCLICommandMap()
	for _, command := range commandMap {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}

func commandExit(config *config, args ...string) error {
	fmt.Print("\nClosing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandMap(config *config, args ...string) error {
	url := ""
	if config.Next == "" {
		url = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	} else {
		url = config.Next
	}
	locationAreas, err := api.GetLocations(url)
	if err != nil {
		return err
	}
	config.Next = locationAreas.Next
	config.Previous = locationAreas.Previous
	for _, result := range locationAreas.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func commandMapB(config *config, args ...string) error {
	url := ""
	if config.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	} else {
		url = config.Previous
	}
	locationAreas, err := api.GetLocations(url)
	if err != nil {
		return err
	}
	config.Next = locationAreas.Next
	config.Previous = locationAreas.Previous
	for _, result := range locationAreas.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func commandExplore(config *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("explore command requires an argument")
	}
	location := strings.ToLower(args[0])
	locationAreaDetails, err := api.GetLocationDetails(location)
	if err != nil {
		return err
	}
	for _, encounter := range locationAreaDetails.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}

func commandCatch(config *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("explore command requires an argument")
	}
	pokemonName := strings.ToLower(args[0])
	pokemon, err := api.GetPokemonDetails(pokemonName)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonName)
	baseExp := pokemon.BaseExperience
	// 36 is the lowest possible base xp and the check will result in
	// 35/36 chance to catch the pokemon with lowest base xp
	// while it will result in 35/608 chance to catch the pokemon with
	// with highest base xp of 608
	if rand.Intn(baseExp) <= 35 {
		fmt.Printf("Gotcha! %v was caught!\n", pokemonName)
		pokedex[pokemonName] = pokemon
	} else {
		fmt.Printf("Oh no! %v got away!\n", pokemonName)
	}
	return nil
}
