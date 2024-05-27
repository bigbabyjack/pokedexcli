package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/bigbabyjack/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*string, *config) error
}

type Repl struct {
	scanner  *bufio.Scanner
	commands map[string]cliCommand
	cache    *pokecache.Cache
}

func newRepl() Repl {
	scanner := bufio.NewScanner(os.Stdin)
	repl := Repl{
		scanner:  scanner,
		commands: make(map[string]cliCommand),
		cache:    pokecache.NewCache(time.Minute * 5),
	}
	return repl
}

func runRepl(cfg *config) {

	repl := newRepl()
	repl.commands = getCommands()
	for {
		fmt.Print("Pokedex > ")
		if !repl.scanner.Scan() {
			break
		}
		input := repl.scanner.Text()
		inputSplit := strings.Split(input, " ")
		if cmd, exists := repl.commands[inputSplit[0]]; exists {
			var arg string
			if len(inputSplit) > 1 {
				arg = inputSplit[1]
			}
			if err := cmd.callback(&arg, cfg); err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
