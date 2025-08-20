package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("username required")
	}
	uerr := s.config.SetUser(cmd.args[0])
	if uerr != nil {
		return uerr
	}
	fmt.Printf("User has been set as: %+v\n", cmd.args[0])
	return nil
}

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
