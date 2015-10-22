package cli_test

import (
	"testing"
	
	"github.com/michaebeam/cli"
)

func TestCli(t *testing.T) {
	args := {"say", "hello", "michael"}

	c := cli.New()
	c.Handle("say", say)
	c.Handle("goodbye", goodbye)
	c.Execute(args)
}

var say := cli.NewCmd()
say.HelpMsg = "say hello|goodbye [name]"
say.ExecFunc = func(args []string) {
	if args == nil {
		fmt.Println(say.Help())
		return
	}

	c := cli.New()
	c.Handle("hello", hello)
	c.Handle("goodbye", goodbye)
	c.HandleDefault(hello)
	c.Execute(args)
}

var hello := cli.NewCmd()
hello.HelpMsg = "hello [name]"
hello.ExecFunc = func(args []string) {
	if args == nil {
		args = append(args, "World")
	}
	fmt.Println("Hello", args, "!")
}

var goodbye := cli.NewCmd()
goodbye.HelpMsg = "goodbye [Name]"
goodbye.ExecFunc = func(args []string) {
	if args == nil {
		args = append(args, "Everybody")
	}
	fmt.Println("Goodbye", args, "!")
}
