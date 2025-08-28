package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/benedekpal/gator/internal/database"
	"github.com/google/uuid"
)

func ToNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{String: "", Valid: false}
	}
	return sql.NullString{String: s, Valid: true}
}

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

func isDuplicateURLError(err error) bool {
	if err == nil {
		return false
	}
	if strings.Contains(err.Error(), "duplicate") {
		return true
	}
	return false
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
		formattedTime, err := time.Parse(time.RFC3339, item.PubDate)
		if err != nil {
			formattedTime, err = time.Parse(time.RFC1123, item.PubDate)
			if err != nil {
				formattedTime = time.Now()
			}
		}

		newPost := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: ToNullString(item.Description),
			PublishedAt: formattedTime,
			FeedID:      feedToFetch.ID,
		}
		_, perr := s.db.CreatePost(context.Background(), newPost)

		if isDuplicateURLError(perr) {
			continue
		}
		if perr != nil {
			// Log other types of errors so you know about them
			log.Printf("Unexpected error saving post '%s': %v", newPost.Title, perr)
		}

	}

	return nil
}
