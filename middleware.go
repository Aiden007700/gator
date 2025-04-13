package main

import (
	"context"
	"errors"
	"github.com/aiden007700/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return errors.New("user not logged in or does not exist")
		}

		return handler(s, cmd, user)
	}
}
