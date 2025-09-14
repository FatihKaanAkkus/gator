package main

import (
	"fmt"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	registered map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if handler, exists := c.registered[cmd.Name]; exists {
		return handler(s, cmd)
	}
	return fmt.Errorf("unknown command: %s", cmd.Name)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registered[name] = f
}
