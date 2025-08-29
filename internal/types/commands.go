package types

import "fmt"

type Command struct {
	Name string
	Args []string
}

func NewCommand(name string, args []string) *Command {
	return &Command{
		Name: name,
		Args: args,
	}
}

type Commands struct {
	Commands map[string]func(*State, Command) error
}

func NewCommands() *Commands {
	return &Commands{
		Commands: make(map[string]func(*State, Command) error),
	}
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.Commands[name] = f
}

func (c *Commands) Run(s *State, cmd *Command) error {
	f, exists := c.Commands[cmd.Name]
	if !exists {
		return fmt.Errorf("unknown command: %s", cmd.Name)

	}

	return f(s, *cmd)
}
