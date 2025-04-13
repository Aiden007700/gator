package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/aiden007700/gator/internal/database"
	"github.com/google/uuid"
	"github.com/mmcdole/gofeed"
	"log"
	"time"
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
func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Printf("Error fetching next feed: %v", err)
		return
	}

	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Error marking feed as fetched: %v", err)
		return
	}

	fmt.Printf("Fetching feed: %s\n", feed.Url)
	fp := gofeed.NewParser()
	parsedFeed, err := fp.ParseURL(feed.Url)
	if err != nil {
		log.Printf("Error parsing feed: %v", err)
		return
	}

	for _, item := range parsedFeed.Items {
		publishedAt, _ := time.Parse(time.RFC1123Z, item.Published) // Handle parsing errors gracefully
		post := database.CreatePostParams{
			ID:          uuid.New(),
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: item.Description != ""},
			PublishedAt: sql.NullTime{Time: publishedAt, Valid: !publishedAt.IsZero()},
			FeedID:      feed.ID,
		}

		err := s.db.CreatePost(context.Background(), post)
		if err != nil {
			log.Printf("Error saving post: %v", err)
		}
	}
}
