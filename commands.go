package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/benedekpal/gator/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("username required")
	}

	dbUser, dberror := s.db.GetUser(context.Background(), cmd.args[0])
	if dberror != nil {
		return dberror
	}

	if dbUser.Name != cmd.args[0] {
		return errors.New("username not registered")
	}

	uerr := s.config.SetUser(cmd.args[0])
	if uerr != nil {
		return uerr
	}

	fmt.Printf("User has been set as: %+v\n", cmd.args[0])
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("username required")
	}

	userinstance := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	}

	dbUser, dberror := s.db.CreateUser(context.Background(), userinstance)

	if err, ok := dberror.(*pq.Error); ok {
		if err.Code == "23505" {
			return fmt.Errorf("username %s already exists", cmd.args[0])
		}
	}

	if dberror != nil {
		return fmt.Errorf("u-u not good %s", dberror)
	}

	s.config.SetUser(dbUser.Name)

	fmt.Printf("User was created\n")
	fmt.Printf("User data:\n")
	fmt.Printf("%+v\n%+v\n%+v\n%+v\n", dbUser.ID, dbUser.CreatedAt, dbUser.UpdatedAt, dbUser.Name)

	return nil
}

func handlerClearUsers(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return errors.New("no arg required")
	}

	err := s.db.ClearUsers(context.Background())

	if err != nil {
		return err
	}

	fmt.Printf("User table was cleared\n")

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
