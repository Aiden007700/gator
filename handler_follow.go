package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/aiden007700/gator/internal/database"
	"github.com/google/uuid"
	"time"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: %v <feed url>", cmd.Name)
	}
	url := cmd.Args[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		if err == sql.ErrNoRows {
			err = handlerAddfeed(s, cmd, user)
			if err != nil {
				return err
			}
			feed, err = s.db.GetFeedByUrl(context.Background(), url)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	createFeedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), createFeedFollowParams)
	if err != nil {
		return err
	}

	fmt.Printf("%+v", feedFollow)

	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	feedFollowsForUser, err := s.db.GetFeedFollowsForUser(context.Background(), user.Name)
	if err != nil {
		return err
	}

	for _, v := range feedFollowsForUser {
		fmt.Printf("- %v", v.FeedName)
	}

	return nil
}

func handleUnFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: %v <feed url>", cmd.Name)
	}
	url := cmd.Args[0]

	err := s.db.DeleteFeedFollowByUserAndUrl(context.Background(), database.DeleteFeedFollowByUserAndUrlParams{
		UserID: user.ID,
		Url:    url,
	})
	if err != nil {
		return err
	}
	return nil
}
