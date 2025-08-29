package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/NachoGz/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func handleRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("there are no arguments, one is expected")
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	})
	if err != nil {
		log.Println("error creating user")
		return err
	}

	s.cfg.SetUser(user.Name)
	log.Println("the user was created succesfully")
	printUser(user)
	return nil
}

// pretty-print for user
func printUser(u database.User) {
	fmt.Println("=============================================================")
	fmt.Println("id:", u.ID)
	fmt.Println("created_at:", u.CreatedAt)
	fmt.Println("updated_at:", u.UpdatedAt)
	fmt.Println("name:", u.Name)
	fmt.Println("=============================================================")
}
