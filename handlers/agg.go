package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/NachoGz/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func handleAgg(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("there are no arguments, one is expected")
	}
	timeBetweenRequests, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		log.Println("couln't parse time_between_reqs")
		return err
	}

	fmt.Println("Collecting feeds every", timeBetweenRequests)
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) error {
	nextFeedToFetch, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Println("there is no next feed")
		return err
	}

	err = s.db.MarkFeedFetched(context.Background(), nextFeedToFetch.ID)
	if err != nil {
		log.Println("couldn't update feed")
		return err
	}

	feed, err := fetchFeed(context.Background(), nextFeedToFetch.Url)
	if err != nil {
		log.Println("error fetching feed")
		return err
	}

	// printFeed(feed)
	for _, item := range feed.Channel.Item {
		s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: parsePubDate(item.PubDate),
			FeedID:      uuid.NullUUID{UUID: nextFeedToFetch.ID, Valid: true},
		})
	}

	return nil
}

func parsePubDate(pubDate string) time.Time {
	parsedTime, err := time.Parse(time.RFC1123Z, pubDate)
	if err != nil {
		log.Println("error parsing publication date:", err)
		return time.Time{}
	}
	return parsedTime
}
