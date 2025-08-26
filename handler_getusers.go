package main

import (
	"context"
	"fmt"
)

func handlerGetUsers(s *state, _ command) error {
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
