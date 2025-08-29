package handlers

import (
	"context"
	"errors"
	"github.com/NachoGz/blog-aggregator/internal/types"
	"log"
)

func HandleReset(s *types.State, cmd types.Command) error {
	if len(cmd.Args) != 0 {
		return errors.New("no arguments are expected")
	}

	err := s.DB.DeleteAllUsers(context.Background())
	if err != nil {
		log.Println("error deleting users")
		return err
	}

	return nil
}
