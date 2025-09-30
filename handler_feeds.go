package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Samuel-Tarifa/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command,u database.User) error {
	if len(cmd.arguments) < 2 {
		return fmt.Errorf("addfeed requires 2 args: name,url")
	}

	name := cmd.arguments[0]
	url := cmd.arguments[1]

	ctx := context.Background()

	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Url:       url,
		Name:      name,
		UserID:    u.ID,
	}

	feed, err := s.db.CreateFeed(ctx, params)

	if err != nil {
		return fmt.Errorf("error creating feed:\n%v", err)
	}

	followParams:=database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: u.ID,
		FeedID: feed.ID,
	}

	_,err=s.db.CreateFeedFollow(ctx,followParams)

	if err!=nil{
		return fmt.Errorf("error creating feedFollow:\n%v",err)
	}

	fmt.Println("Feed:", feed.ID, feed.CreatedAt, feed.UpdatedAt, feed.Name, feed.Url, feed.UserID)

	return nil
}


func handlerFeeds(s *state,cmd command) error {
	
	ctx:=context.Background()

	feeds,err:=s.db.GetFeeds(ctx)

	if err!=nil{
		return fmt.Errorf("error getting feeds:\n%v",err)
	}

	for _,feed:=range feeds{

		fmt.Printf("Feed name: %s\n",feed.Name)
		fmt.Printf("Feed url: %s\n",feed.Url)
		fmt.Printf("User name: %s\n\n",feed.UserName)

	}

	return nil
}