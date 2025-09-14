package main

import (
	"context"
	"fmt"
	"time"

	"github.com/FatihKaanAkkus/gator/internal/database"
	"github.com/google/uuid"
)

func handleAddFeed(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("title is required")
	}
	if len(cmd.Args) < 2 {
		return fmt.Errorf("url is required")
	}
	feedTitle := cmd.Args[0]
	feedUrl := cmd.Args[1]

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("user is logged out")
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedTitle,
		Url:       feedUrl,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to create feed: %v", err)
	}

	fmt.Println("Feed created:")
	printFeed(&feed)
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
