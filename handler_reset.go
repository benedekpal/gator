package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, _ command) error {
	err := s.db.ClearUsers(context.Background())

	if err != nil {
		return err
	}

	fmt.Printf("User table was cleared\n")

	err = s.db.ClearFeeds(context.Background())

	if err != nil {
		return err
	}

	fmt.Printf("Feeds table was cleared\n")

	return nil
}
