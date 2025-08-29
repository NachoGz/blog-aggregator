package handlers

import (
	"context"
	"errors"
	"fmt"
	"github.com/NachoGz/blog-aggregator/internal/types"
	"log"
	"time"

	"github.com/NachoGz/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func HandleRegister(s *types.State, cmd types.Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("there are no arguments, one is expected")
	}

	user, err := s.DB.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	})
	if err != nil {
		log.Println("error creating user")
		return err
	}

	err = s.Cfg.SetUser(user.Name)
	if err != nil {
		return err
	}

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
