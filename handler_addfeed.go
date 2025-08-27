package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/benedekpal/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 2 {
		return errors.New("feed name and url are required")
	}

	feedInstance := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    user.ID,
	}

	feed, dberror := s.db.CreateFeed(context.Background(), feedInstance)

	if dberror != nil {
		// optionally: detect unique violation if using pq
		return dberror
	}

	feedFollowInstance := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    feedInstance.UserID,
		FeedID:    feedInstance.ID,
	}

	_, fferr := s.db.CreateFeedFollow(context.Background(), feedFollowInstance)

	if fferr != nil {
		// optionally: detect unique violation if using pq
		return fferr
	}

	fmt.Println("ID:", feed.ID)
	fmt.Println("CreatedAt:", feed.CreatedAt)
	fmt.Println("UpdatedAt:", feed.UpdatedAt)
	fmt.Println("Name:", feed.Name)
	fmt.Println("URL:", feed.Url)
	fmt.Println("UserID:", feed.UserID)

	return nil
}
