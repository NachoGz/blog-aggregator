package main

import (
	"context"
	"log"

	"github.com/NachoGz/blog-aggregator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		current_user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			log.Println("no user is logged in", current_user.Name)
			return err
		}

		return handler(s, cmd, current_user)
	}
}
