package main

import (
	"context"
	"database/sql"
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

	user, err := s.queries.GetUser(context.Background(), currentUserName)
	if err != nil {
		return fmt.Errorf("Failed to get current user: %w", err)
	}

	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}

	feed, err := s.queries.CreateFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("Failed to create feed: %w", err)
	}

	fmt.Printf("%+v\n", feed)

	return nil
}
