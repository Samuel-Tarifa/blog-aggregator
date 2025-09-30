package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Samuel-Tarifa/blog-aggregator/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func scrapeFeeds(s *state) error {
	ctx := context.Background()
	nextFeed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return fmt.Errorf("error getting nextFeedToFetch:\n%v", err)
	}
	params := database.MarkFeedFetchedParams{
		ID: nextFeed.ID,
		LastFetchedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	feed, err := s.db.MarkFeedFetched(ctx, params)

	if err != nil {
		return fmt.Errorf("error marking feed fetched:\n%v", err)
	}

	RSSFeed, err := fetchFeed(ctx, feed.Url)

	if err != nil {
		return fmt.Errorf("error fetching feed:\n%v", err)
	}

	fmt.Printf("Feed fetched:%s\n\n\n", RSSFeed.Channel.Title)

	for _, item := range RSSFeed.Channel.Item {

		var published sql.NullTime

		layouts := []string{
			time.RFC1123Z, // "Mon, 02 Jan 2006 15:04:05 -0700"
			time.RFC1123,  // "Mon, 02 Jan 2006 15:04:05 MST"
			time.RFC822Z,  // "02 Jan 06 15:04 -0700"
			time.RFC822,   // "02 Jan 06 15:04 MST"
			time.RFC3339,  // "2006-01-02T15:04:05Z07:00"
		}

		for _, layout := range layouts {
			if t, err := time.Parse(layout, item.PubDate); err == nil {
				published = sql.NullTime{Time: t, Valid: true}
				break
			}
		}

		postParams := database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title: sql.NullString{
				String: item.Title,
				Valid:  true,
			},
			Url: item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid:  true,
			},
			FeedID:      feed.ID,
			PublishedAt: published,
		}
		post, err := s.db.CreatePost(ctx, postParams)

		if err != nil{
			var pqErr *pq.Error
			if errors.As(err, &pqErr) && pqErr.Code == "23505"{
				//ignore
			} else {
				fmt.Printf("error creating post:\n%v", err)
			}
		}

		fmt.Printf("post saved: %v\n", post.Title)
	}

	fmt.Printf("\n\n")

	return nil
}
