package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/FatihKaanAkkus/gator/internal/config"
	"github.com/FatihKaanAkkus/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v\n", err)
	}

	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		log.Fatalf("error connecting to db: %v\n", err)
	}
	defer db.Close()

	dbQueries := database.New(db)

	s := &state{
		cfg: &cfg,
		db:  dbQueries,
	}

	cmds := &commands{
		registered: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handleRegister)
	cmds.register("reset", handleReset) // dev only
	cmds.register("users", handleListUsers)
	cmds.register("agg", handleAgg)
	cmds.register("feeds", middlewareLoggedIn(handleListFeeds))
	cmds.register("addfeed", middlewareLoggedIn(handleAddFeed))
	cmds.register("follow", middlewareLoggedIn(handleFollow))
	cmds.register("following", middlewareLoggedIn(handleFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handleUnfollow))
	cmds.register("browse", middlewareLoggedIn(handleBrowse))

	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("Usage: gator <command> [args...]")
	}

	err = cmds.run(s, command{
		Name: args[0],
		Args: args[1:],
	})
	if err != nil {
		log.Fatal(err)
	}
}
