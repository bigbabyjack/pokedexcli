package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

type Repl struct {
	scanner  *bufio.Scanner
	commands map[string]cliCommand
	config   Config
}

type Config struct {
	Next     string
	Previous string
}

func newRepl() Repl {
	scanner := bufio.NewScanner(os.Stdin)
	repl := Repl{
		scanner:  scanner,
		commands: make(map[string]cliCommand),
	}
	repl.getCommands()

	return repl
}

func runRepl() {

	repl := newRepl()
	config := Config{
		Next:     "",
		Previous: "",
	}
	for {
		fmt.Print("Pokedex > ")
		if !repl.scanner.Scan() {
			break
		}
		input := repl.scanner.Text()

		if cmd, exists := repl.commands[input]; exists {
			if err := cmd.callback(&config); err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
