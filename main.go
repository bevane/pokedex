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
			fmt.Print("Unknown command")
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
	fmt.Print("\nWelcome to the Pokedex!\nUsage:\n\n")
	commandMap := getCLICommandMap()
	for _, command := range commandMap {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}

func commandExit() error {
	fmt.Print("\nClosing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func cleanInput(text string) []string {
	trimmedText := strings.TrimRight(strings.TrimLeft(text, " "), " ")
	words := strings.Fields(strings.ToLower(trimmedText))
	return words
}
