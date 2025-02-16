package main

import (
	"context"
	"errors"
	"fmt"
	"log"
)

func handleLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("there are no arguments, one is expected")
	}

	user, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		log.Println("error fetching user")
		return err
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		fmt.Printf("error setting user: %v", err)
		return err
	}

	fmt.Printf("username %s has been set\n", user.Name)
	return nil
}
