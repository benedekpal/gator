package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/benedekpal/gator/internal/database"
)

func handlerBrowse(s *state, cmd command) error {
	var limit int32

	if len(cmd.args) > 1 {
		return errors.New("more parameters were provided, only give  limit (optional)")
	}

	if len(cmd.args) == 0 {
		limit = 2
	} else {
		n, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return err
		}
		limit = int32(n)
	}

	userQuery := database.GetPostsForUserParams{
		Name:  s.config.CurrentUserName,
		Limit: limit,
	}

	userPosts, err := s.db.GetPostsForUser(context.Background(), userQuery)
	if err != nil {
		return err
	}

	for _, p := range userPosts {
		fmt.Printf("%s\t%s\n", p.Title, p.PublishedAt.Format(time.RFC3339))

		desc := "(no description)"
		if p.Description.Valid { // only Description is nullable
			s := p.Description.String
			if len(s) > 120 {
				s = s[:120] + "..."
			}
			desc = s
		}
		fmt.Println(desc)
		fmt.Println()
	}

	return nil
}
