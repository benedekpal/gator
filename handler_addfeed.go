package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/benedekpal/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return errors.New("feed name and url are required")
	}

	currentUser, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)

	if err != nil {
		return err
	}

	feedInstance := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    currentUser.ID,
	}

	feed, dberror := s.db.CreateFeed(context.Background(), feedInstance)

	if dberror != nil {
		// optionally: detect unique violation if using pq
		return dberror
	}

	fmt.Println("ID:", feed.ID)
	fmt.Println("CreatedAt:", feed.CreatedAt)
	fmt.Println("UpdatedAt:", feed.UpdatedAt)
	fmt.Println("Name:", feed.Name)
	fmt.Println("URL:", feed.Url)
	fmt.Println("UserID:", feed.UserID)

	return nil
}
