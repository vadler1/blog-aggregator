package main

import (
	"context"
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
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      username,
	}

	user, err := s.queries.CreateUser(context.Background(), params)

	if err != nil {
		return fmt.Errorf("Failed to register user: %w", err)
	}

	err = s.config.SetUser(user.Name)

	if err != nil {
		return fmt.Errorf("Failed to set user: %w", err)
	}

	fmt.Println("User created successfully:")
	printUser(user)

	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.queries.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Failed to get users: %w", err)
	}

	for _, user := range users {
		suffix := ""
		if user.Name == s.config.CurrentUserName {
			suffix = " (current)"
		}
		fmt.Printf("%s%s\n", user.Name, suffix)
	}

	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}
