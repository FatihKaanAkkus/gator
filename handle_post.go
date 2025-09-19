package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/FatihKaanAkkus/gator/internal/database"
)

func handleBrowse(s *state, cmd command, user database.User) error {
	limit := 10
	if len(cmd.Args) > 0 {
		if val, err := strconv.Atoi(cmd.Args[0]); err == nil {
			limit = val
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("failed to show posts for user: %v", err)
	}

	if len(posts) == 0 {
		fmt.Printf("No posts found for %s.\n", user.Name)
		return nil
	}

	fmt.Printf("Browsing %d posts for user %s (max %d):\n", len(posts), user.Name, limit)
	for _, post := range posts {
		fmt.Println("----------")
		fmt.Printf("* %s from %s\n", post.PublishedAt.Time.Format("Mon Jan 2"), post.FeedName)
		fmt.Printf("* Title: %s\n", post.Title)
		fmt.Printf("    %v\n", post.Description.String)
		fmt.Printf("* Link: %s\n", post.Url)
	}

	return nil
}
