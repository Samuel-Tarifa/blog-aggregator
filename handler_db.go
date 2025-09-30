package main

import (
	"context"
	"fmt"
)


func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error reseting the database:\n%v", err)
	}
	fmt.Printf("Database succesfully reseted")
	return nil
}