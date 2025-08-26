package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/benedekpal/gator/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("username required")
	}

	userinstance := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	}

	dbUser, dberror := s.db.CreateUser(context.Background(), userinstance)

	if err, ok := dberror.(*pq.Error); ok {
		if err.Code == "23505" {
			return fmt.Errorf("username %s already exists", cmd.args[0])
		}
	}

	if dberror != nil {
		return fmt.Errorf("u-u not good %s", dberror)
	}

	s.config.SetUser(dbUser.Name)

	fmt.Printf("User was created\n")
	fmt.Printf("User data:\n")
	fmt.Printf("%+v\n%+v\n%+v\n%+v\n", dbUser.ID, dbUser.CreatedAt, dbUser.UpdatedAt, dbUser.Name)

	return nil
}
