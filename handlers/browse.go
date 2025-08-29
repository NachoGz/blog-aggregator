package handlers

import (
	"context"
	"fmt"
	"github.com/NachoGz/blog-aggregator/internal/types"
	"log"

	"strconv"

	"github.com/NachoGz/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func HandleBrowse(s *types.State, cmd types.Command, current_user database.User) error {
	var limit int32
	if len(cmd.Args) == 0 {
		limit = 2
	} else {
		arg, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return err
		}
		limit = int32(arg)
	}

	posts, err := s.DB.GetPostsFromUser(context.Background(), database.GetPostsFromUserParams{
		UserID: uuid.NullUUID{UUID: current_user.ID, Valid: true},
		Limit:  limit,
	})
	if err != nil {
		log.Println("couldn't create post for user", current_user.Name)
		return err
	}

	for _, post := range posts {
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
		fmt.Println("\tTitle:", post.Title)
		fmt.Println("\tLink:", post.Url)
		fmt.Println("\tDescription:", post.Description)
		fmt.Println("\tPubDate:", post.PublishedAt)
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
	}

	return nil
}
