package main

import (
	"net/http"
	"io"
	"encoding/xml"
	"context"
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
	var v RSSFeed
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "gator")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = xml.Unmarshal(body, &v)
	if err != nil {
		return nil, err
	}

	v.Channel.Title = html.UnescapeString(v.Channel.Title)
	v.Channel.Description = html.UnescapeString(v.Channel.Description)

	for i := range v.Channel.Item {
		v.Channel.Item[i].Title = html.UnescapeString(v.Channel.Item[i].Title) 
		v.Channel.Item[i].Description = html.UnescapeString(v.Channel.Item[i].Description)
	}
 
	return &v, nil
}
