package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/laurafauxvaux/blog_aggregator/internal/database"
	"github.com/lib/pq"
	"github.com/lib/pq/pqerror"
)

func scrapeFeeds(s *state) error {
	ctx := context.Background()

	nextFeed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return err
	}

	if err := s.db.MarkFeedFetched(ctx, nextFeed.ID); err != nil {
		return err
	}
	fmt.Println("Fetching:", nextFeed.Url)
	feedFetched, err := fetchFeed(ctx, nextFeed.Url)
	if err != nil {
		return err
	}

	layouts := []string{
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
	}

	for _, item := range feedFetched.Channel.Item {
		pubDate := sql.NullTime{}

		for _, layout := range layouts {
			parsedDate, err := time.Parse(layout, item.PubDate)
			if err == nil {
				pubDate = sql.NullTime{
					Time:  parsedDate,
					Valid: true,
				}
				break
			}
		}

		pubDesc := sql.NullString{
			String: item.Description,
			Valid:  item.Description != "",
		}

		postParams := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: pubDesc,
			PublishedAt: pubDate,
			FeedID:      nextFeed.ID,
		}

		err := s.db.CreatePost(ctx, postParams)
		pqErr := pq.As(err, pqerror.UniqueViolation)
		if pqErr != nil {
			continue
		}
		if err != nil {
			log.Println("error saving post:", err)
		}
	}
	return nil
}
