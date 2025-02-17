package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/NachoGz/blog-aggregator/internal/config"
	"github.com/NachoGz/blog-aggregator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	cmds map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmds[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	f, exists := c.cmds[cmd.name]
	if !exists {
		return fmt.Errorf("unknown command: %s", cmd.name)
	}

	return f(s, cmd)
}

func main() {
	var s state
	var cmds commands

	cmds.cmds = make(map[string]func(*state, command) error)

	s.cfg = config.Read()

	db, err := sql.Open("postgres", s.cfg.DbUrl)
	if err != nil {
		fmt.Println("error opening the db:", err)
		os.Exit(1)
	}

	s.db = database.New(db)

	cliArgs := os.Args

	if len(cliArgs) < 2 {
		fmt.Println("there must be more than 2 arguments")
		os.Exit(1)
	}

	cmd := command{
		name: cliArgs[1],
		args: cliArgs[2:],
	}

	run := true

	switch cmd.name {
	case "login":
		cmds.register(cmd.name, handleLogin)
	case "register":
		cmds.register(cmd.name, handleRegister)
	case "reset":
		cmds.register(cmd.name, handleReset)
	case "users":
		cmds.register(cmd.name, handlerUsers)
	case "agg":
		cmds.register(cmd.name, handleAgg)
	case "addfeed":
		cmds.register(cmd.name, middlewareLoggedIn(handleAddFeed))
	case "feeds":
		cmds.register(cmd.name, handlerFeeds)
	case "follow":
		cmds.register(cmd.name, middlewareLoggedIn(handleFollow))
	case "following":
		cmds.register(cmd.name, middlewareLoggedIn(handleFollowing))
	case "unfollow":
		cmds.register(cmd.name, middlewareLoggedIn(handleUnfollow))
	case "browse":
		cmds.register(cmd.name, middlewareLoggedIn(handleBrowse))
	case "help":
		run = false
		printUsage()
	}

	if run {
		err = cmds.run(&s, cmd)
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
	fmt.Println("\t reset	\t resets the state of the database, i.e deletes all users and records associated to them. Usage: ./gator reset")
	fmt.Println("\t agg	\t fetches the RSS feeds, parse them and stores the as posts in the database. Usage: ./gator agg <time-between-reqs>")
	fmt.Println("\t addfeed\t adds feeds to the database, usage: ./gator addfeed <feed-name> <feed-url>")
	fmt.Println("\t feeds	\t lists all the feeds in the database, usage: ./gator feeds")
	fmt.Println("\t follow	\t the current user starts to follow the given feed, usage: ./gator follow <feed-url>")
	fmt.Println("\t following\t lists all the feeds that the current user is following, usage: ./gator following")
	fmt.Println("\t unfollow\t the current users unfollows the given feed, usage: ./gator unfollow <feed-url>")
	fmt.Println("\t browse	\t view all the posts from the feeds the user follows, usage: ./gator browse <limit> (default is 2)")
}
