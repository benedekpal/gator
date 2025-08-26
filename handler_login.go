package main

import (
	"context"
	"errors"
	"fmt"
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
