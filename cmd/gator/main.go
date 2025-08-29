package main

import (
	"database/sql"
	"fmt"
	"github.com/NachoGz/blog-aggregator/internal/middleware"
	"github.com/NachoGz/blog-aggregator/internal/types"
	"os"

	"github.com/NachoGz/blog-aggregator/handlers"
	"github.com/NachoGz/blog-aggregator/internal/config"
	"github.com/NachoGz/blog-aggregator/internal/database"
	_ "github.com/lib/pq"
)

func main() {

	s := types.NewState(nil, config.Read())
	db, err := sql.Open("postgres", s.Cfg.DbUrl)
	if err != nil {
		fmt.Println("error opening the db:", err)
		os.Exit(1)
	}
	s.DB = database.New(db)

	commands := types.NewCommands()
	cliArgs := os.Args

	if len(cliArgs) < 2 {
		fmt.Println("there must be more than 2 arguments")
		os.Exit(1)
	}

	cmd := types.NewCommand(cliArgs[1], cliArgs[2:])

	run := true

	switch cmd.Name {
	case "login":
		commands.Register(cmd.Name, handlers.HandleLogin)
	case "register":
		commands.Register(cmd.Name, handlers.HandleRegister)
	case "reset":
		commands.Register(cmd.Name, handlers.HandleReset)
	case "users":
		commands.Register(cmd.Name, handlers.HandlerUsers)
	case "agg":
		commands.Register(cmd.Name, handlers.HandleAgg)
	case "addfeed":
		commands.Register(cmd.Name, middleware.LoggedIn(handlers.HandleAddFeed))
	case "feeds":
		commands.Register(cmd.Name, handlers.HandleFeeds)
	case "follow":
		commands.Register(cmd.Name, middleware.LoggedIn(handlers.HandleFollow))
	case "following":
		commands.Register(cmd.Name, middleware.LoggedIn(handlers.HandleFollowing))
	case "unfollow":
		commands.Register(cmd.Name, middleware.LoggedIn(handlers.HandleUnfollow))
	case "browse":
		commands.Register(cmd.Name, middleware.LoggedIn(handlers.HandleBrowse))
	case "help":
		run = false
		printUsage()
	}

	if run {
		err = commands.Run(s, cmd)
		if err != nil {
			fmt.Println("error running the command: ", err)
			fmt.Println("Run './gator help' for usage.")
			os.Exit(1)
		}
	}

}

func printUsage() {
	fmt.Println("\t\t\t\t\t🐊 Gator is an RSS feed aggregator in Go! 🐊")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("\t./gator <command> [arguments]")
	fmt.Println()
	fmt.Println("The commands are:")
	fmt.Println()
	fmt.Println("\t login	\t sets the current user, usage: ./gator login <username>")
	fmt.Println("\t register\t adds a new user to the database, usage: ./gator register <username>")
	fmt.Println("\t users	\t lists all the users in the database, usage: ./gator users")
	fmt.Println("\t reset	\t resets the types.State of the database, i.e deletes all users and records associated to them. Usage: ./gator reset")
	fmt.Println("\t agg	\t fetches the RSS feeds, parse them and stores the as posts in the database. Usage: ./gator agg <time-between-reqs>")
	fmt.Println("\t addfeed\t adds feeds to the database, usage: ./gator addfeed <feed-name> <feed-url>")
	fmt.Println("\t feeds	\t lists all the feeds in the database, usage: ./gator feeds")
	fmt.Println("\t follow	\t the current user starts to follow the given feed, usage: ./gator follow <feed-url>")
	fmt.Println("\t following\t lists all the feeds that the current user is following, usage: ./gator following")
	fmt.Println("\t unfollow\t the current users unfollows the given feed, usage: ./gator unfollow <feed-url>")
	fmt.Println("\t browse	\t view all the posts from the feeds the user follows, usage: ./gator browse <limit> (default is 2)")
}
