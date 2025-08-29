package middleware

import (
	"context"
	"github.com/NachoGz/blog-aggregator/internal/types"
	"log"

	"github.com/NachoGz/blog-aggregator/internal/database"
)

func LoggedIn(handler func(s *types.State, cmd types.Command, user database.User) error) func(*types.State, types.Command) error {
	return func(s *types.State, cmd types.Command) error {
		currentUser, err := s.DB.GetUser(context.Background(), s.Cfg.CurrentUserName)
		if err != nil {
			log.Println("no user is logged in", currentUser.Name)
			return err
		}

		return handler(s, cmd, currentUser)
	}
}
