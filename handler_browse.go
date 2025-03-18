package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/burush0/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	var limit int
	if len(cmd.Args) == 1 {
		var err error
		limit, err = strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("couln't convert arg to int: %w", err)
		}
	} else {
		limit = 2
	}

	getParams := database.GetPostsByUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}

	posts, err := s.db.GetPostsByUser(context.Background(), getParams)
	if err != nil {
		return fmt.Errorf("couldn't get posts: %w", err)
	}

	for _, post := range posts {
		fmt.Printf("- '%s'\n", post.Title)
	}

	return nil
}
