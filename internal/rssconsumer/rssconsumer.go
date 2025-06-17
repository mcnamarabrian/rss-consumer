package rssconsumer

import (
	"encoding/xml"
	"errors"
	"io"
	"log"
	"net/http"
	"time"
)

// RSS represents an RSS feed structure.
type RSS struct {
	Channel struct {
		Items []Item `xml:"item"`
	} `xml:"channel"`
}

// Item represents an item in the RSS feed.
type Item struct {
	Title   string `xml:"title"`
	Link    string `xml:"link"`
	PubDate string `xml:"pubDate"`
}

// GetItemsSince fetches and parses the RSS feed and returns titles published after the given time.
func GetItemsSince(rssURL string, since time.Time) ([]string, error) {
	resp, err := http.Get(rssURL)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("error closing response body: %v", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch rss feed")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rss RSS
	if err := xml.Unmarshal(body, &rss); err != nil {
		return nil, err
	}

	var result []string
	for _, item := range rss.Channel.Items {
		published, err := parsePubDate(item.PubDate)
		if err != nil {
			continue
		}
		if published.After(since) {
			result = append(result, item.Title)
		}
	}

	return result, nil
}

// parsePubDate attempts to parse various RSS pubDate formats.
func parsePubDate(dateStr string) (time.Time, error) {
	layouts := []string{
		time.RFC1123Z,
		time.RFC1123,
		time.RFC822Z,
		time.RFC822,
		time.RFC3339,
	}

	for _, layout := range layouts {
		if t, err := time.Parse(layout, dateStr); err == nil {
			return t, nil
		}
	}
	return time.Time{}, errors.New("unable to parse date: " + dateStr)
}
