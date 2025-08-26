package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, _ command) error {
	rss, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")

	if err != nil {
		return err
	}

	fmt.Printf("RSSFeed:\n%+v\n", rss)

	return nil
}
