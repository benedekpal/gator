package main

import (
	"context"
	"fmt"

	"github.com/benedekpal/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {

	currentUser := user.Name
	userFeed, err := s.db.GetFeedFollowsForUser(context.Background(), currentUser)

	if err != nil {
		return err
	}

	fmt.Println("Feed-Followed by ", currentUser)

	for _, feed := range userFeed {

		fmt.Println(feed.FeedName)
	}
	fmt.Println("")

	return nil
}
