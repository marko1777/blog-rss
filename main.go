package main

import (
	"fmt"
	"os"

	"github.com/marko1777/blog-rss/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands map[string]func(*state, command) error

func (this commands) register(name string, f func(*state, command) error) {
	this[name] = f
}

func (this commands) run(s *state, cmd command) error {
	command, ok := this[cmd.name]
	if !ok {
		return fmt.Errorf("Command: %s; not found", cmd.name)
	}

	return command(s, cmd)
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Usage: ./blog-rss login <username>")
	}
	username := cmd.args[0]
	err := s.cfg.SetUser(username)

	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("User: %s; has been set\n", username)
	return nil
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Usage <command> <arguments>")
		os.Exit(1)
	}

	state := &state{
		cfg: config.Read(),
	}

	cmds := commands{}
	cmds.register("login", handlerLogin)

	cmd := args[0]
	cmdArgs := args[1:]
	err := cmds.run(state, command{
		name: cmd,
		args: cmdArgs,
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
