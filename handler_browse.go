package main

import (
	"context"
	"fmt"
	"github.com/aiden007700/gator/internal/database"
	"strconv"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.Args) > 0 {
		var err error
		limit, err = strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("invalid limit: %v", err)
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Printf("Title: %s\nURL: %s\nPublished At: %v\n\n", post.Title, post.Url, post.PublishedAt.Time)
	}

	return nil
}
