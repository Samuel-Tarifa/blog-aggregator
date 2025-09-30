package main

import (
	"context"
	"fmt"

	"github.com/Samuel-Tarifa/blog-aggregator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error{
	
	return func(s *state, c command) error {
		u,err:=s.db.GetUser(context.Background(),s.cfg.CurrentUserName)
		if err!=nil{
			return fmt.Errorf("error getting user:\n%v",err)
		}
		return handler(s,c,u)
	}
	
}