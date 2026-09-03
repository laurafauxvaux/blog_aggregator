package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("no argument required")
	}

	follows, err := s.db.GetFeedFollowsForUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	for _, follow := range follows {
		fmt.Println(follow.FeedNames)
	}

	return nil
}

// GetFeedFollowsForUser(ctx context.Context, name string) ([]GetFeedFollowsForUserRow, error)
