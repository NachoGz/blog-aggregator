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
		cmds.register(cmd.name, handleAddFeed)
	case "feeds":
		cmds.register(cmd.name, handlerFeeds)
	}

	err = cmds.run(&s, cmd)
	if err != nil {
		fmt.Println("error running the command: ", err)
		os.Exit(1)
	}
}
