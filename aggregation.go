package main

import (
	"context"
	"fmt"
	"log"

	"github.com/burush0/gator/internal/database"
)

// func scrapeFeeds(s *state) error {
// 	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
// 	if err != nil {
// 		return fmt.Errorf("couldn't fetch next feed from db: %w", err)
// 	}

// 	err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
// 	if err != nil {
// 		return fmt.Errorf("couldn't mark feed as fetched: %w", err)
// 	}

// 	feed, err := fetchFeed(context.Background(), nextFeed.Url)
// 	if err != nil {
// 		return fmt.Errorf("couldn't fetch feed: %w", err)
// 	}

// 	fmt.Printf("%+v\n", feed.Channel.Title)
// 	return nil
// }

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Println("Couldn't get next feeds to fetch", err)
		return
	}
	log.Println("Found a feed to fetch!")
	scrapeFeed(s.db, feed)
}

func scrapeFeed(db *database.Queries, feed database.Feed) {
	err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Couldn't mark feed %s fetched: %v", feed.Name, err)
		return
	}

	feedData, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("Couldn't collect feed %s: %v", feed.Name, err)
		return
	}
	for _, item := range feedData.Channel.Item {
		fmt.Printf("Found post: %s\n", item.Title)
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(feedData.Channel.Item))
}
