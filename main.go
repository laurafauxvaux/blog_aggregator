package main

import (
	"fmt"
	"log"

	"github.com/laurafauxvaux/blog_aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("failed to read initial config file: %v", err)
	}

	if err := cfg.SetUser("laura"); err != nil {
		log.Fatalf("failed to set new user: %v", err)
	}

	updatedCfg, err := config.Read()
	if err != nil {
		log.Fatalf("failed to read updated config file: %v", err)
	}

	fmt.Println(updatedCfg)
}
