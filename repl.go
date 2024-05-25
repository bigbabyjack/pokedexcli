package main

import (
	"bufio"
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
