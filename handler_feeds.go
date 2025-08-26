package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, _ command) error {
	feeds, err := s.db.GetFeeds(context.Background())

	if err != nil {
		return err
	}

	fmt.Println("")

	for i, feed := range feeds {
		fmt.Println("Feed ", i+1)
		fmt.Println(" ID:\t", feed.Name)
		fmt.Println(" Url:\t", feed.Url)
		fmt.Println(" UserName:\t", feed.Name_2)
		fmt.Println("")
	}

	return nil
}
