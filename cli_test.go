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
	c.HandleDefaultFunc(echo)
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
	c.HandleDefaultFunc(echo)
	c.Execute(args[1:])
}

func echo(args []string) {
	fmt.Println(strings.Join(args, " "))
}

func hello(args []string) {
	if len(args) <= 1 {
		args = []string{"World"}
	} else {
		args = args[1:]
	}
	fmt.Printf("Hello %s!\n", strings.Join(args, " "))
}

func goodbye(args []string) {
	if len(args) <= 1 {
		args = []string{"Everybody"}
	} else {
		args = args[1:]
	}
	fmt.Printf("Goodbye %s!\n", strings.Join(args, " "))
}
