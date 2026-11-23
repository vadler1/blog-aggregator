package main

import (
	"context"
	"fmt"

	"github.com/vadler1/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		if s.config.CurrentUserName == "" {
			return fmt.Errorf("User not logged in")
		}
		currentUserName := s.config.CurrentUserName

		user, err := s.queries.GetUserByName(context.Background(), currentUserName)
		if err != nil {
			return fmt.Errorf("Failed to get current user: %w", err)
		}
		return handler(s, cmd, user)
	}
}
