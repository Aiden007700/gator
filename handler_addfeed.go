package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/aiden007700/gator/internal/database"
	"github.com/google/uuid"
	"time"
)

func handlerAddfeed(s *state, cmd command) error {
	if len(cmd.Args) < 2 {
		return errors.New("addfeed expects feedName andurl")
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return err
	}

	fmt.Println(feed)
	return nil
}
