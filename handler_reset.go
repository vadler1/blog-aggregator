package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	// clear users
	err := s.queries.ClearUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Failed to clear users: %w", err)
	}
	fmt.Println("Users cleared successfully")
	return nil
}
