package main

import (
	"fmt"
	"os"

	"database/sql"

	_ "github.com/lib/pq"
	"github.com/vadler1/gator/internal/config"
	"github.com/vadler1/gator/internal/database"
)

type state struct {
	config  *config.Config
	queries *database.Queries
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("Error reading config:", err)
		os.Exit(1)
	}

	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		os.Exit(1)
	}
	defer db.Close()

	queries := database.New(db)

	s := state{
		config:  cfg,
		queries: queries,
	}

	cmds := commands{
		commands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)

	args := os.Args

	if len(args) < 2 {
		fmt.Println("No command provided")
		os.Exit(1)
	}

	cmd := command{
		name: args[1],
		args: args[2:],
	}

	err = cmds.run(&s, cmd)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
