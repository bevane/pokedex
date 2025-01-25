package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bevane/pokedex/internal/api"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	Next     string
	Previous string
}

var pokedex = make(map[string]api.Pokemon)

func startRepl() {
	locationConfig := &config{}
	scanner := bufio.NewScanner(os.Stdin)
	nameToCliCommand := getCLICommandMap()
	for {
		var command cliCommand
		var err error
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		command, ok := nameToCliCommand[cleanedInput[0]]
		if !ok {
			fmt.Print("Unknown command")
			command = nameToCliCommand["help"]
		}
		err = command.callback(locationConfig, cleanedInput[1:]...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func getCLICommandMap() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display names of the next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display names of the previous 20 location areas",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Display list of pokemon in an area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch and add a pokemon to Pokedex",
			callback:    commandCatch,
		},
	}
}

func cleanInput(text string) []string {
	trimmedText := strings.TrimRight(strings.TrimLeft(text, " "), " ")
	words := strings.Fields(strings.ToLower(trimmedText))
	return words
}
