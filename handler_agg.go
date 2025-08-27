package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("time_between_reqs required")
	}

	delay, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf(" Collecting feeds every %+v\n", delay)

	ticker := time.NewTicker(delay)
	for ; ; <-ticker.C {
		err := scrapeFeeds(s)

		if err != nil {
			return err
		}
	}

	return nil
}

func scrapeFeeds(s *state) error {
	feedToFetch, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	rss, err := fetchFeed(context.Background(), feedToFetch.Url)

	if err != nil {
		return err
	}

	err = s.db.MarkFeedFetched(context.Background(), feedToFetch.ID)

	if err != nil {
		return err
	}

	for _, item := range rss.Channel.Item {
		fmt.Printf("RSSTitle:%+v\n", item.Title)
	}

	return nil
}
