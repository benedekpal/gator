package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerGetUsers(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return errors.New("no arg required")
	}

	users, err := s.db.GetUsers(context.Background())

	if err != nil {
		return err
	}

	currentUser := s.config.CurrentUserName

	for _, user := range users {
		fmt.Printf("* ")
		fmt.Printf("%s", user.Name)
		if user.Name == currentUser {
			fmt.Printf(" (current)")
		}
		fmt.Printf("\n")
	}
	return nil
}
