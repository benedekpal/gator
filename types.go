package main

import (
	"github.com/benedekpal/gator/internal/config"
	"github.com/benedekpal/gator/internal/database"
)

type state struct {
	config *config.Config
	db     *database.Queries
}

type command struct {
	name string
	args []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}
