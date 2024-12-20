package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/marko1777/blog-rss/internal/database"
	"github.com/marko1777/blog-rss/internal/state"
)

type Command struct {
	Name string
	Args []string
}

type Commands map[string]func(*state.State, Command) error

func (this Commands) Register(
	Name string,
	f func(*state.State, Command) error,
) {
	this[Name] = f
}

func (this Commands) Run(s *state.State, cmd Command) error {
	Command, ok := this[cmd.Name]
	if !ok {
		return fmt.Errorf("Command: %s; not found", cmd.Name)
	}

	return Command(s, cmd)
}

func HandlerLogin(s *state.State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Usage: ./blog-rss login <username>")
	}
	username := cmd.Args[0]
	_, err := s.DBQueries.GetUser(context.Background(), username)

	if err != nil {
		fmt.Println(err)
		return err
	}
	err = s.Cfg.SetUser(username)

	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("User: %s; has been set\n", username)
	return nil
}

func HandlerRegister(s *state.State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Usage: ./blog-rss register <username>")
	}
	username := cmd.Args[0]
	_, err := s.DBQueries.CreateUser(
		context.Background(),
		database.CreateUserParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      username,
		},
	)

	if err != nil {
		fmt.Println(err)
		return err
	}

	err = s.Cfg.SetUser(username)

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("User: %s; registered\n", username)

	return nil
}

func HandlerReset(s *state.State, cmd Command) error {
	err := s.DBQueries.Reset(context.Background())
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("User table reset")

	return nil
}

func HandlerUsers(s *state.State, cmd Command) error {
	users, err := s.DBQueries.GetUsers(context.Background())
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, user := range users {
		str := "* " + user.Name
		if user.Name == s.Cfg.CurrentUserName {
			str += " (current)"
		}
		fmt.Println(str)
	}
	return nil
}
