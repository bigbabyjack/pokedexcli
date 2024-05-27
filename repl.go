package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/bigbabyjack/pokedexcli/internal/pokecache"
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
	cache    *pokecache.Cache
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
		config: struct {
			Next     string
			Previous string
		}{
			Next:     "",
			Previous: "",
		},
		cache: pokecache.NewCache(time.Minute * 5),
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
			if err := cmd.callback(&repl.config); err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
