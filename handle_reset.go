package main

import (
	"context"
	"fmt"
)

func handleReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to delete all users: %v", err)
	}
	fmt.Println("Database reset successfully!")
	return nil
}
