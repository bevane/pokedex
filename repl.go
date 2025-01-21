package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Next     string
	Previous string
}

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
		command, ok := nameToCliCommand[input]
		if !ok {
			fmt.Print("Unknown command")
			command = nameToCliCommand["help"]
		}
		err = command.callback(locationConfig)
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
	}
}

func cleanInput(text string) []string {
	trimmedText := strings.TrimRight(strings.TrimLeft(text, " "), " ")
	words := strings.Fields(strings.ToLower(trimmedText))
	return words
}
