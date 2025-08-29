package handlers

import (
	"context"
	"errors"
	"fmt"
	"github.com/NachoGz/blog-aggregator/internal/types"
	"log"
	"time"

	"github.com/NachoGz/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func HandleFollow(s *types.State, cmd types.Command, currentUser database.User) error {
	if len(cmd.Args) == 0 {
		return errors.New("there are no arguments, one is expected")
	}

	feed, err := s.DB.GetFeedFromURL(context.Background(), cmd.Args[0])
	if err != nil {
		log.Println("error fetching feed")
		return err
	}

	_, err = s.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    uuid.NullUUID{UUID: currentUser.ID, Valid: true},
		FeedID:    uuid.NullUUID{UUID: feed.ID, Valid: true},
	})
	if err != nil {
		log.Println("error creating feed follow")
		return err
	}

	fmt.Printf("feed %s followed by %s\n", feed.Name, currentUser)

	return nil
}

func HandleFollowing(s *types.State, cmd types.Command, current_user database.User) error {
	if len(cmd.Args) != 0 {
		return errors.New("there are no arguments, one is expected")
	}

	feedsFollowed, err := s.DB.GetFeedFollowsForUser(context.Background(), uuid.NullUUID{
		UUID:  current_user.ID,
		Valid: true,
	})
	if err != nil {
		log.Println("error fetching followed feeds")
		return err
	}

	fmt.Println("Feeds followed by", current_user.Name)
	for _, feedFollow := range feedsFollowed {
		fmt.Println("* ", feedFollow.FeedName)
	}

	return nil
}

func HandleUnfollow(s *types.State, cmd types.Command, currentUser database.User) error {
	if len(cmd.Args) == 0 {
		return errors.New("there are no arguments, one is expected")
	}
	err := s.DB.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		Name: currentUser.Name,
		Url:  cmd.Args[0],
	})
	if err != nil {
		log.Println("couldn't unfollow", cmd.Args[0])
	}

	return nil
}
