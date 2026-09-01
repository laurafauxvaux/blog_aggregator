package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/laurafauxvaux/blog_aggregator/internal/config"
	"github.com/laurafauxvaux/blog_aggregator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	dbQueries := database.New(db)

	newState := state{
		db:  dbQueries,
		cfg: &cfg,
	}

	cmds := commands{
		names: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerListUsers)

	if len(os.Args) < 2 {
		log.Fatalf("not enough arguments")
	}
	cmd := command{
		name:      os.Args[1],
		arguments: os.Args[2:],
	}

	if err := cmds.run(&newState, cmd); err != nil {
		log.Fatalf("error: %v", err)
	}

}
