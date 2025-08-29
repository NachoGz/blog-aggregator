package handlers

import (
	"context"
	"errors"
	"fmt"
	"github.com/NachoGz/blog-aggregator/internal/types"
	"log"
)

func HandlerUsers(s *types.State, cmd types.Command) error {
	if len(cmd.Args) != 0 {
		return errors.New("no arguments are expected")
	}

	users, err := s.DB.GetUsers(context.Background())
	if err != nil {
		log.Println("error fetching users")
		return err
	}

	for _, user := range users {
		if user.Name == s.Cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Println("* ", user.Name)
		}
	}

	return nil
}
