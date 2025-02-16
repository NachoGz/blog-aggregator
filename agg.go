package main

import (
	"context"
	"errors"
	"log"
)

func handleAgg(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("there are no arguments, one is expected")
	}

	feed, err := fetchFeed(context.Background(), cmd.args[0])
	if err != nil {
		log.Println("error fetching feed")
		return err
	}

	printFeed(feed)

	return nil
}
