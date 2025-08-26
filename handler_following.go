package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, _ command) error {

	currentUser := s.config.CurrentUserName
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
