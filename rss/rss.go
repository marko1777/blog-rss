package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

type Feed struct {
	Channel struct {
		Title       string `xml:"title"`
		Link        string `xml:"link"`
		Description string `xml:"description"`
		Item        []Item `xml:"item"`
	} `xml:"channel"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func FetchFeed(ctx context.Context, feedURL string) (*Feed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	req.Header.Set("User-Agent", "blog-rss")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	feed := &Feed{}
	if err := xml.Unmarshal(respBytes, feed); err != nil {
		fmt.Println(err)
		return nil, err
	}

	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)

	for i, item := range feed.Channel.Item {
		feed.Channel.Item[i].Title = html.UnescapeString(item.Title)
		feed.Channel.Item[i].Description = html.UnescapeString(item.Description)
	}

	return feed, nil
}
