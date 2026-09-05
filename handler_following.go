package main

import (
	"context"
	"fmt"

	"github.com/laurafauxvaux/blog_aggregator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("no argument required")
	}

	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.Name)
	if err != nil {
		return err
	}

	for _, follow := range follows {
		fmt.Println(follow.FeedNames)
	}

	return nil
}
