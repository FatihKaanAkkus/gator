package main

import (
	"context"
	"fmt"
	"time"

	"github.com/FatihKaanAkkus/gator/internal/database"
	"github.com/google/uuid"
)

func handleFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("feed url is required")
	}
	feedUrl := cmd.Args[0]

	feed, err := s.db.GetFeed(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("failed to get feed: %v", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to follow feed: %v", err)
	}

	fmt.Println("Feed follow created:")
	printFeedFollow(feedUrl, &feedFollow)
	return nil
}

func handleFollowing(s *state, cmd command, user database.User) error {
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("failed to get followed feed for user: %v", err)
	}

	fmt.Println("Following feeds:")
	for _, feedFollow := range feedFollows {
		fmt.Printf("* %s (%s) (%s)\n", feedFollow.FeedID, feedFollow.FeedName, feedFollow.UserName)
	}

	return nil
}

func handleUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("feed url is required")
	}
	feedUrl := cmd.Args[0]

	feed, err := s.db.GetFeed(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("failed to get feed: %v", err)
	}

	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to unfollow: %v", err)
	}

	return nil
}

func printFeedFollow(url string, feedFollow *database.CreateFeedFollowRow) {
	fmt.Printf("* URL:      %s\n", url)
	fmt.Printf("* Feed:     %s\n", feedFollow.FeedName)
	fmt.Printf("* User:     %s\n", feedFollow.UserName)
}
