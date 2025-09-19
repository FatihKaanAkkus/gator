package main

import (
	"context"
	"fmt"
	"time"

	"github.com/FatihKaanAkkus/gator/internal/database"
)

func handleAgg(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("time between requests is required")
	}
	timeBetweenReqs := cmd.Args[0]

	duration, err := time.ParseDuration(timeBetweenReqs)
	if err != nil {
		return fmt.Errorf("invalid duration: %v", err)
	}

	fmt.Printf("Collecting feeds every %v\n", duration)

	ticker := time.NewTicker(duration)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return
	}

	fmt.Println("-----------")
	fmt.Printf("Fetching next feed: %s\n", feed.Url)
	printFeed(&feed)

	scrapeFeed(s.db, feed)
}

func scrapeFeed(db *database.Queries, feed database.Feed) {
	err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		fmt.Printf("cannot mark feed as fetched: %v", err)
		return
	}

	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		fmt.Printf("cannot fetch rss feed: %v", err)
	}

	fmt.Println("* Posts found:")
	for _, item := range rssFeed.Channel.Item {
		if item.Title != "" {
			fmt.Printf("  |-- %s\n", item.Title)
		}
	}

	fmt.Printf("* Fetched %v posts from %s.", len(rssFeed.Channel.Item), feed.Name)
}
