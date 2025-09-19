package main

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/FatihKaanAkkus/gator/internal/database"
	"github.com/google/uuid"
)

func handleAgg(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: %v <time between requests>", cmd.Name)
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
		fmt.Printf("could not mark feed as fetched: %v", err)
		return
	}

	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		fmt.Printf("could not fetch rss feed: %v", err)
	}

	fmt.Printf("* Posts found: %d\n", len(rssFeed.Channel.Item))

	countCreates := 0
	for _, item := range rssFeed.Channel.Item {
		created, err := createPostFromFeed(db, feed, item)
		if err != nil {
			fmt.Printf("  |-- %s (%v)\n", item.Title, err)
		}
		if created {
			countCreates++
		}
	}

	fmt.Printf("* Fetched %d new posts from %s.", countCreates, feed.Name)
}

func createPostFromFeed(db *database.Queries, feed database.Feed, item RSSItem) (bool, error) {
	publishedAt := sql.NullTime{}
	if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
		publishedAt = sql.NullTime{
			Time:  t,
			Valid: true,
		}
	}

	post, err := db.CreatePost(context.Background(), database.CreatePostParams{
		ID:          uuid.New(),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		Url:         item.Link,
		Title:       item.Title,
		Description: sql.NullString{Valid: true, String: item.Description},
		PublishedAt: publishedAt,
		FeedID:      feed.ID,
	})
	if err != nil {
		if strings.Contains(err.Error(), "unique constraint \"posts_url_key\"") {
			return false, nil
		} else {
			return true, fmt.Errorf("could not create post: %w", err)
		}
	}

	fmt.Printf("  |++ %s (created)\n", post.Title)
	return true, nil
}
