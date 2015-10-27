// Copyright 2015 Michael Beam
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

/*
Package cli provides basic tools for creating hireactical command line interfaces.

	// ...
	flag.Parse()
	
	c := cli.New()
	c.HelpMsg = "cmd start|stop"
	c.HandleFunc("start", start)
	c.HandleFunc("stop", stop)
	c.HandleDefault(start)
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

// Objects implementing the Command interface can be registered to handle
// paticular sub-commands in a cli.
type Command interface {
	Execute([]string)
}

// Cli is a command interface. It executes the subcommand from its arguments
// list or a default command.
type Cli struct {
	commands map[string]Command
	defaultCmd Command
	HelpMsg string
}

// New Creates a new Cli to be used to handle a set of sub-commands.
func New() *Cli {
	return &Cli{commands: make(map[string]Command)}
}

// Execute checks the first argument and executes the matching sub-command if
// any. If no arguments are present it executes the Cli.defaultCmd. If the
// first argument does not match a sub-command it prints the Cli.HelpMsg.
func (c *Cli) Execute(args []string) {
	if len(args) < 1 {
		if c.defaultCmd != nil {
			c.defaultCmd.Execute(args)
		} else {
			fmt.Println(c.HelpMsg)
		}
	} else if cmd, ok := c.commands[args[0]]; ok {
		cmd.Execute(args[1:])
	} else {
		fmt.Println(c.HelpMsg)
	}
}

// Handle registers the handler for a given command.
func (c *Cli) Handle(cmd string, handler Command) {
	c.commands[cmd] = handler
}

// HandleDefault registers the default command to be executed if no arguments
// are given. If no arguments are given and Cli.defaultCmd == nil the
// Cli.HelpMsg will be printed.
func (c *Cli) HandleDefault(handler Command) {
	c.defaultCmd = handler
}

// HandleDefaultFunc registers a function to handle the default command to be
// executed if no arguments are given.
func (c *Cli) HandleDefaultFunc(handler cmdFunc) {
	c.defaultCmd = handler
}

// HandleFunc registers a function to handle a given command.
func (c *Cli) HandleFunc(cmd string, handler cmdFunc) {
	c.commands[cmd] = handler
}

// cmdFunc allows a func([]string) to be cast as a handler that supports the
// Command interface.
type cmdFunc func([]string)

func (c cmdFunc) Execute(args []string) {
	c(args)
}
