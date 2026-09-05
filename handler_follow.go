package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/laurafauxvaux/blog_aggregator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("URL is required")
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.arguments[0])
	if err != nil {
		return err
	}

	followParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	follow, err := s.db.CreateFeedFollow(context.Background(), followParams)
	if err != nil {
		return err
	}

	fmt.Println(follow.UserName, follow.FeedName)

	return nil
}
