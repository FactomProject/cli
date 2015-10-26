// Copyright 2015 Michael Beam
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

/*
Package cli provides basic tools for creating hireactical command line interfaces.

	// ...
	flag.Parse()
	
	c := cli.New()
	c.HelpMsg = "cmd start|stop"
	c.HandleFunc("start", start, "cmd start")
	c.HandleFunc("stop", stop, "cmd stop")
	c.Execute(flag.Args())
	
	func start(args []string) {
		// Start the server
	}
	
	func stop(args []string) {
		// Stop the server
	}
*/

package cli

import (
	"fmt"
)

// Objects implementing the Command interface can be registered to handle paticular sub-commands in a cli.
type Command interface {
	Execute([]string)
}

type Cli struct {
	commands map[string]Command
	defaultCmd Command
	HelpMsg string
}

func New() *Cli {
	return &Cli{commands: make(map[string]Command)}
}

func (c *Cli) Execute(args []string) {
	if len(args) < 1 {
		if c.defaultCmd != nil {
			c.defaultCmd.Execute(args)
		} else {
			fmt.Println(c.Help())
		}
	} else if cmd, ok := c.commands[args[0]]; ok {
		cmd.Execute(args[1:])
	} else {
		fmt.Println(c.HelpMsg)
	}
}

func (c *Cli) Handle(cmd string, handler Command) {
	c.commands[cmd] = handler
}

func (c *Cli) HandleDefault(handler Command) {
	c.defaultCmd = handler
}

func (c *Cli) HandleDefaultFunc(handler cmdFunc) {
	c.defaultCmd = handler
}

func (c *Cli) HandleFunc(cmd string, handler cmdFunc) {
	c.commands[cmd] = handler
}

func (c *Cli) Help() string {
	return c.HelpMsg
}

type cmdFunc func([]string)

func (c cmdFunc) Execute(args []string) {
	c(args)
}
