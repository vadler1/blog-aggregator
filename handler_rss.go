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

func handlerAddfeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("Expected 2 arguments, got %d", len(cmd.args))
	}

	name := cmd.args[0]
	url := cmd.args[1]

	createFeedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}

	feed, err := s.queries.CreateFeed(context.Background(), createFeedParams)
	if err != nil {
		return fmt.Errorf("Failed to create feed: %w", err)
	}

	followFeedParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	feedFollow, err := s.queries.CreateFeedFollow(context.Background(), followFeedParams)
	if err != nil {
		return fmt.Errorf("Failed to create feed follow: %w", err)
	}

	fmt.Println("Feed created successfully:")
	printFeed(feed)
	fmt.Println()
	fmt.Println("Feed followed successfully:")
	printFeedFollow(feedFollow.UserName, feedFollow.FeedName)
	fmt.Println("=====================================")

	return nil
}

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Expected 1 argument, got %d", len(cmd.args))
	}
	feedURL := cmd.args[0]

	feed, err := s.queries.GetFeedByUrl(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("Failed to get feed by URL: %w", err)
	}

	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	feedFollow, err := s.queries.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("Failed to create feed follow: %w", err)
	}

	fmt.Println("Feed follow created successfully:")
	printFeedFollow(feedFollow.UserName, feedFollow.FeedName)

	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("Expected 0 arguments, got %d", len(cmd.args))
	}

	followed_feeds, err := s.queries.GetFeedFollowForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("Failed to retrieve followed feeds for user: %w", err)
	}

	if len(followed_feeds) == 0 {
		fmt.Println("No feed follows found for this user.")
		return nil
	}

	fmt.Printf("Feed follows for user %s:\n", user.Name)
	for _, followed_feed := range followed_feeds {
		fmt.Printf("* %s\n", followed_feed.FeedName)
	}

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

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Expected 1 argument, got %d", len(cmd.args))
	}
	feedUrl := cmd.args[0]

	feed, err := s.queries.GetFeedByUrl(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("Failed to get feed by URL: %w", err)
	}

	err = s.queries.DeleteFeedFollowByFeedId(context.Background(), feed.ID)
	if err != nil {
		return fmt.Errorf("Failed to delete feed follow: %w", err)
	}

	fmt.Println("Feed follow deleted successfully.")

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

func printFeedFollow(username, feedname string) {
	fmt.Printf("* User:          %s\n", username)
	fmt.Printf("* Feed:          %s\n", feedname)
}
