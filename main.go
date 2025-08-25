package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/benedekpal/gator/internal/config"
	"github.com/benedekpal/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	s := state{}
	cmds := commands{handlers: make(map[string]func(*state, command) error)}

	cnfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	fmt.Printf("Read config: %+v\n", cnfg)

	db, err := sql.Open("postgres", cnfg.DBURL)
	if err != nil {
		log.Fatalf("error opening sql: %v", err)
	}

	dbQueries := database.New(db)

	s.db = dbQueries
	s.config = &cnfg

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerClearUsers)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	commandName := os.Args[1]
	commandArgs := os.Args[2:]
	cmd := command{
		name: commandName,
		args: commandArgs,
	}

	err = cmds.run(&s, cmd)

	if err != nil {
		log.Fatal(err)
	}

	cnfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config again: %+v\n", cnfg)

}
