package main

import (
	"log"
	"os"

	"github.com/laurafauxvaux/blog_aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}

	newState := state{
		cfg: &cfg,
	}

	cmds := commands{
		names: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)

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
