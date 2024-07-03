package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	for {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")
	scanner.Scan()
	text := scanner.Text()

	cleaned := cleanInput(text)
	if len(cleaned) == 0 {
		continue
	}

	commandName := cleaned[0]

	availableCommands := getCommands()

	command, ok := availableCommands[commandName]
	if !ok {
		fmt.Println("Invalid command")
		continue
	}
	err := command.callback(cfg)
	if err != nil {
		fmt.Println(err)
	}
	}	
}

type cliCommand struct {
	name string
	description string
	callback func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Prints the help menu",
			callback: callbackHelp,
		},
		"exit": {
			name: "exit",
			description: "Turns off the Pokedex",
			callback: callbackExit,
		},
		"map": {
			name: "map",
			description: "Lists some location areas",
			callback: callbackMap,
		},
		"mapb": {
			name: "mapb",
			description: "Lists the previous location areas",
			callback: callbackMapb,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}