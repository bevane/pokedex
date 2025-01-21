package main

import (
	"fmt"
	"os"
)

func commandHelp(*config) error {
	fmt.Print("\nWelcome to the Pokedex!\nUsage:\n\n")
	commandMap := getCLICommandMap()
	for _, command := range commandMap {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}

func commandExit(*config) error {
	fmt.Print("\nClosing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandMap(*config) error {
	return nil
}
