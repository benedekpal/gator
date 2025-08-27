package main

import (
	"context"
	"errors"

	"github.com/benedekpal/gator/internal/database"
)

func handlerUnFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return errors.New("url required")
	}

	deleteFeedFollow := database.DeleteFeedFollowForUserParams{
		Name: user.Name,
		Url:  cmd.args[0],
	}

	err := s.db.DeleteFeedFollowForUser(context.Background(), deleteFeedFollow)

	if err != nil {
		return err
	}

	return nil
}
