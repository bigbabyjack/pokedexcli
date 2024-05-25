package main

func (r *Repl) getCommands() {
	r.commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    r.commandHelp,
	}

	r.commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    r.commandExit,
	}

	r.commands["map"] = cliCommand{
		name:        "map",
		description: "Display locations.",
		callback:    r.commandMap,
	}

	r.commands["mapb"] = cliCommand{
		name:        "mapb",
		description: "Display previous locations.",
		callback:    r.commandMapb,
	}

}
