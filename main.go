package main

import (
	"log"
	"os"

	"github.com/FatihKaanAkkus/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v\n", err)
	}

	s := &state{
		cfg: &cfg,
	}

	cmds := &commands{
		registered: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	err = cmds.run(s, command{
		Name: args[0],
		Args: args[1:],
	})
	if err != nil {
		log.Fatal(err)
	}
}
