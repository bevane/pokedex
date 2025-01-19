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
	callback    func() error
}

func main() {
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
			fmt.Printf("%s is not a valid command", input)
			command = nameToCliCommand["help"]
		}
		err = command.callback()
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
	}
}

func commandHelp() error {
	fmt.Print("\nWelcome to the Pokedex.\nUsage:\n\n")
	commandMap := getCLICommandMap()
	for _, command := range commandMap {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func cleanInput(text string) []string {
	out := []string{}
	trimmedText := strings.TrimRight(strings.TrimLeft(text, " "), " ")
	words := strings.Split(trimmedText, " ")
	for _, word := range words {
		if word != " " && word != "" {
			out = append(out, word)
		}
	}

	return out
}
