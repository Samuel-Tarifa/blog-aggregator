package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Samuel-Tarifa/blog-aggregator/internal/config"
	"github.com/Samuel-Tarifa/blog-aggregator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	var currState state
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	currState.cfg = &cfg

	dbURL:=currState.cfg.DbURL

	db, err := sql.Open("postgres", dbURL)

	if err!=nil{
		fmt.Printf("Error opnening database: %v\n",err)
		os.Exit(1)
	}

	dbQueries:=database.New(db)

	currState.db=dbQueries

	var currCommands commands
	currCommands.commandMap = map[string]func(*state, command) error{}

	currCommands.register("login", handlerLogin)
	currCommands.register("register",handlerRegister)
	currCommands.register("reset",handlerReset)
	currCommands.register("users",handlerListUsers)
	currCommands.register("agg",handlerAgg)
	currCommands.register("addfeed",middlewareLoggedIn(handlerAddFeed))
	currCommands.register("feeds",handlerFeeds)
	currCommands.register("follow",middlewareLoggedIn(handlerFollow))
	currCommands.register("following",middlewareLoggedIn(handlerFollowing))
	currCommands.register("unfollow",middlewareLoggedIn(handlerUnfollow))
	currCommands.register("browse",middlewareLoggedIn(handlerBrowse))

	arguments := os.Args

	if len(arguments) < 2 {
		fmt.Printf("No commands typed\n")
		os.Exit(1)
	}

	cmdName := arguments[1]
	arguments = arguments[2:]
	cmd := command{
		name:      cmdName,
		arguments: arguments,
	}
	if err = currCommands.run(&currState, cmd); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
