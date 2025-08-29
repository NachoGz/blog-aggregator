package handlers

import (
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/NachoGz/blog-aggregator/internal/types"
	"html"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/NachoGz/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	feed := &RSSFeed{}

	// make request to the RSS feed
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		log.Println("error fetching the feed")
		return nil, err
	}

	// send request
	req.Header.Set("User-Agent", "gator")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("error making the request")
		return nil, err
	}

	// read response and unmarshal it to the RSSFeed struct
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("error reading response body")
		return nil, err
	}

	err = xml.Unmarshal(body, feed)
	if err != nil {
		log.Println("error unmarshaling response")
		return nil, err
	}

	// decode escaped HTML entities
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	for i := range feed.Channel.Item {
		feed.Channel.Item[i].Title = html.UnescapeString(feed.Channel.Item[i].Title)
		feed.Channel.Item[i].Description = html.UnescapeString(feed.Channel.Item[i].Description)
	}

	return feed, nil
}

func HandleAddFeed(s *types.State, cmd types.Command, current_user database.User) error {
	name, url := cmd.Args[0], cmd.Args[1]

	feed, err := s.DB.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID: uuid.NullUUID{
			UUID:  current_user.ID,
			Valid: true,
		},
	})
	if err != nil {
		log.Println("error creating feed")
		return err
	}

	_, err = s.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    uuid.NullUUID{UUID: current_user.ID, Valid: true},
		FeedID:    uuid.NullUUID{UUID: feed.ID, Valid: true},
	})
	if err != nil {
		log.Println("error creating feed follow")
		return err
	}

	printDbFeed(feed)

	return nil
}

func HandleFeeds(s *types.State, cmd types.Command) error {
	if len(cmd.Args) != 0 {
		return errors.New("no arguments are expected")
	}

	feeds, err := s.DB.GetFeeds(context.Background())
	if err != nil {
		log.Println("error fetching feeds")
		return err
	}

	for _, feed := range feeds {
		username, err := s.DB.GetUserName(context.Background(), feed.UserID.UUID)
		if err != nil {
			log.Println("erro fetching user", feed.UserID.UUID)
			return err
		}
		fmt.Println("=============================================================")
		fmt.Println("name:", feed.Name)
		fmt.Println("url:", feed.Url)
		fmt.Println("creator:", username)
		fmt.Println("=============================================================")
	}

	return nil
}

func printFeed(feed *RSSFeed) {
	fmt.Println("==============================================================================================================================")
	fmt.Println("Title:", feed.Channel.Title)
	fmt.Println("Link:", feed.Channel.Title)
	fmt.Println("Description:", feed.Channel.Title)
	for _, item := range feed.Channel.Item {
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------")
		fmt.Println("\tTitle:", item.Title)
		fmt.Println("\tLink:", item.Link)
		fmt.Println("\tDescription:", item.Description)
		fmt.Println("\tPubDate:", item.PubDate)
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------")

	}
	fmt.Println("==============================================================================================================================")
}

// pretty-print for feed in db
func printDbFeed(f database.Feed) {
	fmt.Println("=============================================================")
	fmt.Println("id:", f.ID)
	fmt.Println("created_at:", f.CreatedAt)
	fmt.Println("updated_at:", f.UpdatedAt)
	fmt.Println("name:", f.Name)
	fmt.Println("url:", f.Url)
	fmt.Println("user_id:", f.UserID.UUID)
	fmt.Println("=============================================================")
}
