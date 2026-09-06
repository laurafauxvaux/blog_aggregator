package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/laurafauxvaux/blog_aggregator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := int32(2)
	if len(cmd.arguments) > 0 {
		i64, err := strconv.ParseInt(cmd.arguments[0], 10, 32)
		if err != nil {
			return fmt.Errorf("couldn't convert the limit: %w", err)
		}
		limit = int32(i64)
	}

	postParams := database.GetPostForUserParams{
		UserID: user.ID,
		Limit:  limit,
	}

	posts, err := s.db.GetPostForUser(context.Background(), postParams)
	if err != nil {
		return fmt.Errorf("error getting posts:%w", err)
	}

	for _, post := range posts {
		fmt.Println(post.Title)
		fmt.Println(post.Url)
	}

	return nil
}
