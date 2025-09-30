package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Samuel-Tarifa/blog-aggregator/internal/database"
	"github.com/google/uuid"
)


func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("login command expects a single argument, the username")
	}

	name := cmd.arguments[0]
	_, err := s.db.GetUser(context.Background(), name)
	if err == sql.ErrNoRows {
		return fmt.Errorf("user not found")
	}
	if err != nil {
		return err
	}

	if err := s.cfg.SetUser(cmd.arguments[0]); err != nil {
		return err
	}

	fmt.Println("User setted correclty")

	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("register command expects a single argument, the username")
	}

	id := uuid.New()
	created_at := time.Now()
	updated_at := time.Now()
	name := cmd.arguments[0]
	params := database.CreateUserParams{ID: id, CreatedAt: created_at, UpdatedAt: updated_at, Name: name}
	u, err := s.db.CreateUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error creating user:\n%v", err)
	}

	s.cfg.CurrentUserName = u.Name
	fmt.Printf("User created succesfully:\nID: %v\nCreatedAt: %v\nUpdatedAt: %v\nName: %v\n", u.ID, u.CreatedAt, u.UpdatedAt, u.Name)

	if err := s.cfg.SetUser(cmd.arguments[0]); err != nil {
		return err
	}

	return nil
}

func handlerListUsers(s *state,cmd command) error {
	users,err:=s.db.GetUsers(context.Background())

	if err!=nil{
		return fmt.Errorf("error getting users:\n%v",err)
	}

	currUser:=s.cfg.CurrentUserName

	for _,u:=range users{
		fmt.Printf("* %v",u.Name)
		if u.Name==currUser{
			fmt.Print(" (current)")
		}
		fmt.Printf("\n")
	}

	return nil
}