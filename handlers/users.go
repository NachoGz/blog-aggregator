package main

import (
	"context"
	"errors"
	"fmt"
	"log"
)

func handlerUsers(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return errors.New("no arguments are expected")
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		log.Println("error fetching users")
		return err
	}

	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Println("* ", user.Name)
		}
	}

	return nil
}
