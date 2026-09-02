package main

import (
	"context"
	"fmt"
)

const URL = "https://www.wagslane.dev/index.xml"

func handlerAgg(_ *state, _ command) error {
	feed, err := fetchFeed(context.Background(), URL)
	if err != nil {
		return err
	}
	fmt.Println(feed)
	return nil
}
