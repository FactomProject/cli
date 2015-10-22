package cli

import (
	"fmt"
)

type Command interface {
	Execute([]string)
	Help() string
}

type Cli struct {
	commands map[string]Command
	defaultCmd Command
}

func New() *Cli {
	return &Cli{ commands: make(map[string]Command) }
}

func (c *Cli) Execute(args []string) {
	if len(args) < 1 {
		c.defaultCmd.Execute(args)
	}
	if cmd, ok := c.commands[args[0]]; ok {
		cmd.Execute(args[1:])
	}
	fmt.Println(c.Help())
}

func (c *Cli) Handle(cmd string, handler Command) {
	c.commands[cmd] = handler
}

func (c *Cli) HandleFunc(cmd string, handler func(args []string), help string) {
	a := new(Cmd)
	a.ExecFunc = handler
	a.HelpMsg = help
	c.commands[cmd] = a
}

func (c *Cli) Help() string {
	return "TODO - get help"
}

type Cmd struct {
	ExecFunc func([]string)
	HelpMsg string
}

func NewCmd() *Cmd {
	return new(Cmd)
}

func (c *Cmd) Execute(args []string) {
	c.ExecFunc(args)
}

func (c *Cmd) Help() string {
	return c.HelpMsg
}
