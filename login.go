package main

import (
	"errors"
	"fmt"
)

func handleLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("there are no arguments, one is expected")
	}
	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		fmt.Printf("error setting user: %v", err)
		return err
	}

	fmt.Printf("username %s has been set\n", cmd.args[0])
	return nil
}
