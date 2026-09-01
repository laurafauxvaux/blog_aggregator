package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, _ command) error {
	if err := s.db.DeleteUsers(context.Background()); err != nil {
		return err
	}

	fmt.Println("'Users' has been emptied")

	return nil
}
