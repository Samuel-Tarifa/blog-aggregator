package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Samuel-Tarifa/blog-aggregator/internal/database"
)

func handlerBrowse(s *state, cmd command, u database.User) error {

	var limit int32 = 2

	if len(cmd.arguments) >= 1 {
		n, err := strconv.Atoi(cmd.arguments[0])
		if err != nil {
			fmt.Printf("bad limit input, keeping default: 2")
		} else {
			limit = int32(n)
		}
	}

	ctx := context.Background()

	params := database.GetPostsForUserParams{
		ID:    u.ID,
		Limit: int32(limit),
	}
	posts, err := s.db.GetPostsForUser(ctx, params)
	if err!=nil{
		return fmt.Errorf("error getting posts:\n%v",err)
	}

	for _,post:=range posts{
		fmt.Printf("%s\n%s\n",post.Title.String,post.Description.String)
	}

	return nil
}
