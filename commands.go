package main

import (
	"fmt"
	"github.com/bevane/pokedex/internal/api"
	"os"
)

func commandHelp(config *config) error {
	fmt.Print("\nWelcome to the Pokedex!\nUsage:\n\n")
	commandMap := getCLICommandMap()
	for _, command := range commandMap {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}

func commandExit(config *config) error {
	fmt.Print("\nClosing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandMap(config *config) error {
	url := ""
	if config.Next == "" {
		url = "https://pokeapi.co/api/v2/location-area"
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

func commandMapB(config *config) error {
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
