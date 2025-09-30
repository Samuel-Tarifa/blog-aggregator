package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Samuel-Tarifa/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow (s *state,cmd command,u database.User) error{
	if len(cmd.arguments)<1{
		return fmt.Errorf("follow needs one argument, the feed url")
	}

	url:=cmd.arguments[0]
	ctx:=context.Background()

	feed,err:=s.db.GetFeedByUrl(ctx,url)

	if err!=nil{
		return fmt.Errorf("error getting feed from url:\n%v",err)
	}

	params:=database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: u.ID,
		FeedID: feed.ID,
	}

	createdFeedFollow,err:=s.db.CreateFeedFollow(ctx,params)

	if err!=nil{
		return fmt.Errorf("error creating feedFollow:\n%v",err)
	}

	fmt.Printf("feed followed: %s\n",createdFeedFollow.FeedName)
	fmt.Printf("current user: %s\n",createdFeedFollow.UserName)
	
	return nil
}

func handlerFollowing(s *state,cmd command,u database.User) error {
	
	username:=s.cfg.CurrentUserName
	ctx:=context.Background()

	following,err:=s.db.GetFeedFollowsForUser(ctx,u.ID)
	if err!=nil{
		return fmt.Errorf("error getting feeds folloing:\n%v",err)
	}

	fmt.Printf("Current user: %s\n",username)
	fmt.Printf("Follows:\n\n")
	for _,feed:=range following{
		fmt.Printf("Feed name: %s\n",feed.FeedName)
		fmt.Printf("Feed url: %s\n\n",feed.Url)
	}
	
	return nil
}

func handlerUnfollow(s *state,cmd command,u database.User) error{

	if len(cmd.arguments)<1{
		return fmt.Errorf("unfollow expects 1 argument: feedURL")
	}
	
	ctx:=context.Background()
	url:=cmd.arguments[0]
	
	feed,err:=s.db.GetFeedByUrl(ctx,url)

	if err!=nil{
		return fmt.Errorf("error getting feed by url:\n%v",err)
	}

	params:=database.DeleteFeedFollowParams{
		UserID: u.ID,
		FeedID: feed.ID,
	}
	_,err=s.db.DeleteFeedFollow(ctx,params)
	
	if err!=nil{
		return fmt.Errorf("error deleting FeedFollow:\n%v",err)
	}

	fmt.Printf("User %s unfollowed feed %s\n",u.Name,feed.Name)

	return nil
}