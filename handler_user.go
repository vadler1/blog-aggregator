package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Invalid number of arguments, expected 1")
	}

	username := cmd.args[0]

	err := s.config.SetUser(username)

	if err != nil {
		return fmt.Errorf("Failed to set user: %w", err)
	}

	fmt.Println("User set successfully")

	return nil
}
