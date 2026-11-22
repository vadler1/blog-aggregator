package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/vadler1/gator/internal/database"
	"github.com/vadler1/gator/internal/rss"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("Expected 1 argument, got %d", len(cmd.args))
	}

	// feedURL := cmd.args[0]
	feedURL := "https://www.wagslane.dev/index.xml"
	feed, err := rss.FetchFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("Failed to fetch feed: %w", err)
	}

	fmt.Printf("%+v\n", feed)

	return nil
}

func handlerAddfeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("Expected 2 arguments, got %d", len(cmd.args))
	}

	name := cmd.args[0]
	url := cmd.args[1]

	currentUserName := s.config.CurrentUserName

	user, err := s.queries.GetUserByName(context.Background(), currentUserName)
	if err != nil {
		return fmt.Errorf("Failed to get current user: %w", err)
	}

	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}

	feed, err := s.queries.CreateFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("Failed to create feed: %w", err)
	}

	fmt.Println("Feed created successfully:")
	printFeed(feed)
	fmt.Println()
	fmt.Println("=====================================")

	return nil
}

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.queries.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("Failed to get feeds: %w", err)
	}

	for _, feed := range feeds {

		fmt.Printf("* Name:          %s\n", feed.Name)
		fmt.Printf("* URL:           %s\n", feed.Url)
		user, err := s.queries.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("Failed to get user by ID: %w", err)
		}

		fmt.Printf("* UserName:      %s\n", user.Name)
		fmt.Println()
	}

	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* UserID:        %s\n", feed.UserID)
}
