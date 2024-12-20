package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/marko1777/blog-rss/cmd"
	"github.com/marko1777/blog-rss/internal/config"
	"github.com/marko1777/blog-rss/internal/database"
	"github.com/marko1777/blog-rss/internal/state"

	_ "github.com/lib/pq"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Usage <command> <arguments>")
		os.Exit(1)
	}
	cfg := config.Read()

	db, err := sql.Open("postgres", cfg.DB_URL)

	dbQueries := database.New(db)
	state := &state.State{
		Cfg:       cfg,
		DBQueries: dbQueries,
	}

	cmds := cmd.Commands{}
	cmds.Register("login", cmd.HandlerLogin)
	cmds.Register("register", cmd.HandlerRegister)
	cmds.Register("users", cmd.HandlerUsers)
	cmds.Register("reset", cmd.HandlerReset)
	cmds.Register("agg", cmd.HandlerAgg)
	cmds.Register("addfeed", cmd.HandlerAddFeed)

	cmdName := args[0]
	cmdArgs := args[1:]
	err = cmds.Run(state, cmd.Command{
		Name: cmdName,
		Args: cmdArgs,
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
