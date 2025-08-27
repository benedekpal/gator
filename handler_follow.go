package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/benedekpal/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return errors.New("url required")
	}

	currentFeedInstance, err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	feedFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    currentFeedInstance.ID,
	}

	newFeedFollow, err := s.db.CreateFeedFollow(context.Background(), feedFollow)
	if err != nil {
		return err
	}

	fmt.Println("New Feed-Follow created")
	fmt.Println("Feed: ", newFeedFollow.FeedName)
	fmt.Println("User: ", newFeedFollow.UserName)

	return nil
}
