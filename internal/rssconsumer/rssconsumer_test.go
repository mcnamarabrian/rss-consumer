package rssconsumer

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

func TestGetItemsSince(t *testing.T) {
	data, err := os.ReadFile("testdata/feed.rss")
	if err != nil {
		t.Fatalf("Failed to read test RSS file: %v", err)
	}

	// Start a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(data))
		if err != nil {
			t.Fatalf("Failed to write response: %v", err)
		}
	}))

	defer server.Close()

	since := time.Date(2025, 5, 19, 0, 0, 0, 0, time.UTC)

	titles, err := GetItemsSince(server.URL, since)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(titles) != 1 || !strings.Contains(titles[0], "Release 7.9.0 now available") {
		t.Errorf("Expected 1 new article, got %v", titles)
	}
}
