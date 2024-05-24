package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type Repl struct {
	scanner  *bufio.Scanner
	commands map[string]cliCommand
}

func (r *Repl) commandHelp() error {
	fmt.Println("Welcome to Pokedex!")
	fmt.Println("Usage:")

	if len(r.commands) == 0 {
		return errors.New("Repl has no commands!")
	}

	for c := range r.commands {
		fmt.Println(c, ": ", r.commands[c].description)
	}

	return nil
}

func (r *Repl) commandExit() error {
	fmt.Println("Exiting Pokedex. Goodbye!")
	os.Exit(0)
	return nil
}

func runRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	repl := Repl{
		scanner:  scanner,
		commands: make(map[string]cliCommand),
	}

	repl.commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    repl.commandHelp,
	}

	repl.commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    repl.commandExit,
	}

	for {
		fmt.Print("Pokedex > ")
		if !repl.scanner.Scan() {
			break
		}
		input := repl.scanner.Text()

		if cmd, exists := repl.commands[input]; exists {
			if err := cmd.callback(); err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
