package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: %v <time_between_reqs>", cmd.Name)
	}

	// Parse the time duration
	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %v", err)
	}

	fmt.Printf("Collecting feeds every %v\n", timeBetweenRequests)

	// Set up a ticker and signal handling for graceful shutdown
	ticker := time.NewTicker(timeBetweenRequests)
	defer ticker.Stop()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	// Run the scrapeFeeds function in a loop
	for {
		select {
		case <-ticker.C:
			scrapeFeeds(s)
		case <-quit:
			fmt.Println("Shutting down...")
			return nil
		}
	}
}
