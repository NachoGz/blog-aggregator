package main

import (
	"fmt"
	"os"

	"github.com/NachoGz/blog-aggregator/internal/config"
)

type state struct {
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

	cliArgs := os.Args

	if len(cliArgs) < 2 {
		fmt.Println("there must be more than 2 arguments")
		os.Exit(1)
	}

	cmd := command{
		name: cliArgs[1],
		args: cliArgs[2:],
	}

	cmds.register(cmd.name, handleLogin)

	err := cmds.run(&s, cmd)
	if err != nil {
		fmt.Printf("error running the command: %v", err)
		os.Exit(1)
	}
}
