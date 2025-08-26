package main

import (
	"fmt"
)

func (c *commands) run(s *state, cmd command) error {
	command, exists := c.handlers[cmd.name]
	if !exists {
		return fmt.Errorf("command does not exist %s", cmd.name)
	}

	err := command(s, cmd)
	if err != nil {
		return err
	}

	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}
