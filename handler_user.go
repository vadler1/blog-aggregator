package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/vadler1/gator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Invalid number of arguments, expected 1")
	}

	username := cmd.args[0]

	user, err := s.queries.GetUser(context.Background(), username)

	if err != nil {
		return fmt.Errorf("Failed to get user: %w", err)
	}

	err = s.config.SetUser(user.Name)

	if err != nil {
		return fmt.Errorf("Failed to set user: %w", err)
	}

	fmt.Println("User set successfully")

	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Invalid number of arguments, expected 1")
	}

	username := cmd.args[0]

	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		Name:      username,
	}

	user, err := s.queries.CreateUser(context.Background(), params)

	if err != nil {
		return fmt.Errorf("Failed to register user: %w", err)
	}

	fmt.Println("User registered successfully")

	err = s.config.SetUser(user.Name)

	if err != nil {
		return fmt.Errorf("Failed to set user: %w", err)
	}

	fmt.Printf("User ID: %s\n", user.ID)
	fmt.Printf("Created At: %s\n", user.CreatedAt.Time)
	fmt.Printf("Updated At: %s\n", user.UpdatedAt.Time)
	fmt.Printf("Name: %s\n", user.Name)

	return nil
}

func handlerReset(s *state, cmd command) error {
	// clear users
	err := s.queries.ClearUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Failed to clear users: %w", err)
	}
	fmt.Println("Users cleared successfully")
	return nil
}
