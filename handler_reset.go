package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if err := s.db.DeleteUsers(context.Background()); err != nil {
		fmt.Println("User Reset was NOT successful")
		return err
	}

	fmt.Println("User Reset was successful")
	return nil
}
