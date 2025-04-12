package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeedsWithUser(context.Background())
	if err != nil {
		return err
	}

	for _, v := range feeds {
		fmt.Printf("Name:%v, URL: %v, User: %v \n", v.Name, v.Url, v.Name_2)
	}

	return nil
}
