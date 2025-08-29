package handlers

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/NachoGz/blog-aggregator/internal/types"
)

func HandleLogin(s *types.State, cmd types.Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("there are no arguments, one is expected")
	}

	user, err := s.DB.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		log.Println("error fetching user")
		return err
	}

	err = s.Cfg.SetUser(user.Name)
	if err != nil {
		fmt.Printf("error setting user: %v", err)
		return err
	}

	fmt.Printf("username %s has been set\n", user.Name)
	return nil
}
