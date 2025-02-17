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

func handleFollow(s *state, cmd command, current_user database.User) error {
	if len(cmd.args) == 0 {
		return errors.New("there are no arguments, one is expected")
	}

	feed, err := s.db.GetFeedFromURL(context.Background(), cmd.args[0])
	if err != nil {
		log.Println("error fetching feed")
		return err
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    uuid.NullUUID{UUID: current_user.ID, Valid: true},
		FeedID:    uuid.NullUUID{UUID: feed.ID, Valid: true},
	})
	if err != nil {
		log.Println("error creating feed follow")
		return err
	}

	fmt.Printf("feed %s followed by %s\n", feed.Name, current_user)

	return nil
}

func handleFollowing(s *state, cmd command, current_user database.User) error {
	if len(cmd.args) != 0 {
		return errors.New("there are no arguments, one is expected")
	}

	feeds_followed, err := s.db.GetFeedFollowsForUser(context.Background(), uuid.NullUUID{
		UUID:  current_user.ID,
		Valid: true,
	})
	if err != nil {
		log.Println("error fetching followed feeds")
		return err
	}

	fmt.Println("Feeds followed by", current_user.Name)
	for _, feed_follow := range feeds_followed {
		fmt.Println("* ", feed_follow.FeedName)
	}

	return nil
}

func handleUnfollow(s *state, cmd command, current_user database.User) error {
	if len(cmd.args) == 0 {
		return errors.New("there are no arguments, one is expected")
	}
	err := s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		Name: current_user.Name,
		Url:  cmd.args[0],
	})
	if err != nil {
		log.Println("couldn't unfollow", cmd.args[0])
	}

	return nil
}
