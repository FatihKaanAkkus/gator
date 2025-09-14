package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	fmt.Println(cmd)
	if len(cmd.Args) < 1 {
		return fmt.Errorf("username is required")
	}
	username := cmd.Args[0]

	err := s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("cannot set user: %w", err)
	}

	fmt.Printf("Logged in as %s\n", username)
	return nil
}
