package main

import (
	"context"
	"errors"
	"log"
)

func handleReset(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return errors.New("no arguments are expected")
	}

	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		log.Println("error deleting users")
		return err
	}

	return nil
}
