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
		return fmt.Errorf("explore command requires an argument\n")
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
		return fmt.Errorf("catch command requires an argument\n")
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

func commandInspect(config *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("inspect command requires an argument\n")
	}
	pokemonName := strings.ToLower(args[0])
	pokemon, ok := pokedex[pokemonName]
	if !ok {
		fmt.Printf("You have not caught %v yet\n", pokemonName)
		return nil
	}
	var hp int
	var attack int
	var defense int
	var specialAttack int
	var specialDefense int
	var speed int
	var typesStr string
	for _, stat := range pokemon.Stats {
		switch stat.Stat.Name {
		case "hp":
			hp = stat.BaseStat
		case "attack":
			attack = stat.BaseStat
		case "defense":
			defense = stat.BaseStat
		case "special-attack":
			specialAttack = stat.BaseStat
		case "special-defense":
			specialDefense = stat.BaseStat
		case "speed":
			speed = stat.BaseStat
		}
	}
	for _, pokemonType := range pokemon.Types {
		typesStr += fmt.Sprintf(" - %v\n", pokemonType.Type.Name)
	}

	fmt.Printf(`Name: %v
Height: %v
Weight: %v
Stats:
  -hp: %v
  -attack: %v
  -defense: %v
  -special-attack: %v
  -special-defense: %v
  -speed: %v
Types:
%v`, pokemon.Name, pokemon.Height, pokemon.Weight, hp, attack, defense, specialAttack, specialDefense, speed, typesStr)
	return nil

}

func commandPokedex(config *config, args ...string) error {
	if len(pokedex) == 0 {
		fmt.Println("Pokedex is empty")
		return nil
	}
	pokemonList := "Your Pokedex:\n"
	for _, pokemon := range pokedex {
		pokemonList += fmt.Sprintf(" - %v\n", pokemon.Name)
	}
	fmt.Print(pokemonList)
	return nil
}
