package main

import (
	"fmt"

	"github.com/laurafauxvaux/blog_aggregator/internal/config"
	"github.com/laurafauxvaux/blog_aggregator/internal/database"
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
	names map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if handler, ok := c.names[cmd.name]; !ok {
		return fmt.Errorf("command doesn't exist")
	} else {
		return handler(s, cmd)
	}
}

func (c *commands) register(name string, f func(*state, command) error) {
	if c.names == nil {
		c.names = make(map[string]func(*state, command) error)
	}
	c.names[name] = f
}
