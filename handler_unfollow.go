package main

import (
	"context"
	"fmt"

	"github.com/laurafauxvaux/blog_aggregator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) < 1 {
		return fmt.Errorf("URL is required")
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.arguments[0])
	if err != nil {
		return err
	}

	deleteParams := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	if err := s.db.DeleteFeedFollow(context.Background(), deleteParams); err != nil {
		return err
	}

	fmt.Println("Feed", feed.Name, "has been unfollowed")

	return nil
}
