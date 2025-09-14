package main

import (
	"context"
	"fmt"
	"time"

	"github.com/FatihKaanAkkus/gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("username is required")
	}
	username := cmd.Args[0]

	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("cannot find user: %w", err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("cannot set user: %w", err)
	}

	fmt.Printf("Logged in as %s\n", username)
	return nil
}

func handleRegister(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("username is required")
	}
	username := cmd.Args[0]
	createdAt := time.Now()
	UpdatedAt := time.Now()
	uuid := uuid.New()

	user, err := s.db.GetUser(context.Background(), username)
	if err == nil {
		return fmt.Errorf("user already exists: %v", user.Name)
	}

	user, err = s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid,
		CreatedAt: createdAt,
		UpdatedAt: UpdatedAt,
		Name:      username,
	})
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("could not set current user: %w", err)
	}

	fmt.Println("User created:")
	printUser(&user)
	return nil
}

func handleListUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %v (current)\n", user.Name)
		} else {
			fmt.Printf("* %v\n", user.Name)
		}
	}

	return nil
}

func printUser(user *database.User) {
	fmt.Printf(" * ID:    %v\n", user.ID)
	fmt.Printf(" * Name:  %v\n", user.Name)
}
