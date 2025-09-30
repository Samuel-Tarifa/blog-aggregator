package main

import (
	"fmt"

	"github.com/Samuel-Tarifa/blog-aggregator/internal/config"
	"github.com/Samuel-Tarifa/blog-aggregator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

type command struct {
	name      string
	arguments []string
}

type commands struct {
	commandMap map[string]func(*state, command) error
}

func (c commands) run(s *state, cmd command) error {
	f, ok := c.commandMap[cmd.name]
	if !ok {
		return fmt.Errorf("command doesn't exist")
	}
	if err := f(s, cmd); err != nil {
		return err
	}
	fmt.Printf("Command %s run correctly\n", cmd.name)
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandMap[name] = f
	fmt.Printf("Command %s registered\n", name)
}
