package cli_test

import (
	"fmt"
	"strings"
	"testing"
	
	"github.com/michaelbeam/cli"
)

func TestCli(t *testing.T) {
	args1 := []string{"say", "hello", "Michael"}
	args2 := []string{"goodbye"}
	args3 := []string{}
	args4 := []string{"giberish", "asdf"}
	args5 := []string{"say", "giberish", "asdf"}

	c := cli.New()
	c.HelpMsg = "command say hello|goodbye [name] | command goodbye [name]"
	c.HandleFunc("say", say)
	c.HandleFunc("goodbye", goodbye)
	c.HandleDefaultFunc(hello)
	c.Execute(args1)
	c.Execute(args2)
	c.Execute(args3)
	c.Execute(args4)
	c.Execute(args5)
}

func say(args []string) {
	c := cli.New()
	c.HelpMsg = "say hello|goodbye [name]"
	c.HandleFunc("hello", hello)
	c.HandleFunc("goodbye", goodbye)
	c.HandleDefaultFunc(hello)
	c.Execute(args)
}

func hello(args []string) {
	if len(args) < 1 {
		args = append(args, "World")
	}
	fmt.Printf("Hello %s!\n", strings.Join(args, " "))
}

func goodbye(args []string) {
	if len(args) < 1 {
		args = append(args, "Everybody")
	}
	fmt.Printf("Goodbye %s!\n", strings.Join(args, " "))
}
