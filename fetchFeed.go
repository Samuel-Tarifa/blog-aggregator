package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"html"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)

	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error creating request:\n%v",err)
	}

	cli := http.Client{}

	req.Header.Set("User-Agent", "gator")

	res, err := cli.Do(req)

	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error doing request:\n%v",err)
	}

	if res.StatusCode!=200{
		return &RSSFeed{},fmt.Errorf("status code not 200:\n%v",res.StatusCode)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error reading response:\n%v",err)
	}

	var data RSSFeed

	err = xml.Unmarshal(body, &data)

	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error decoding response:\n%v",err)
	}

	data.Channel.Description=html.UnescapeString(data.Channel.Description)
	data.Channel.Title=html.UnescapeString(data.Channel.Title)

	for i,item := range data.Channel.Item{
		data.Channel.Item[i].Description=html.UnescapeString(item.Description)
		data.Channel.Item[i].Title=html.UnescapeString(item.Title)
	}

	return &data, nil
}
