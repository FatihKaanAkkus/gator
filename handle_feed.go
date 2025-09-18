package main

import (
	"context"
	"fmt"
	"time"

	"github.com/FatihKaanAkkus/gator/internal/database"
	"github.com/google/uuid"
)

func handleAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("title is required")
	}
	if len(cmd.Args) < 2 {
		return fmt.Errorf("url is required")
	}
	feedTitle := cmd.Args[0]
	feedUrl := cmd.Args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      feedTitle,
		Url:       feedUrl,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to create feed: %v", err)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to create feed follow: %v", err)
	}

	// TODO: Rollback when one of the queries fail

	fmt.Println("Feed created and added to following:")
	printFeed(&feed)
	return nil
}

func handleListFeeds(s *state, cmd command, user database.User) error {
	feeds, err := s.db.GetFeedsWithUserName(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get feed for user: %v", err)
	}

	for _, feed := range feeds {
		fmt.Printf("* %s (%s) (%s)\n", feed.Name, feed.Url, feed.Username.String)
	}

	return nil
}

func printFeed(feed *database.Feed) {
	fmt.Printf("* ID:       %s\n", feed.ID)
	fmt.Printf("* Created:  %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:  %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:     %s\n", feed.Name)
	fmt.Printf("* URL:      %s\n", feed.Url)
	fmt.Printf("* UserID:   %s\n", feed.UserID)
}
