package main

import (
	"context"
	"fmt"
)

func handleAgg(s *state, cmd command) error {
	fetchUrl := "https://www.wagslane.dev/index.xml"

	rssFeed, err := fetchFeed(context.Background(), fetchUrl)
	if err != nil {
		return err
	}

	fmt.Println(*rssFeed)

	return nil
}
