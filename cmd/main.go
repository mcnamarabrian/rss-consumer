package main

import (
	"context"
	"errors"
	"log/slog"
	"os"

	"github.com/mcnamarabrian/rssconsumer/internal/rssconsumer"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	Titles []string `json:"titles"`
}

func handler(ctx context.Context, event events.CloudWatchEvent) (Response, error) {
	rssURL := os.Getenv("RSS_URL")
	if rssURL == "" {
		slog.Error("missing RSS_URL environment variable")
		return Response{}, errors.New("RSS_URL environment variable is not set")
	}

	since := event.Time
	slog.Info("processing RSS feed",
		slog.String("rss_url", rssURL),
		slog.Time("since", since),
	)

	titles, err := rssconsumer.GetItemsSince(rssURL, since)
	if err != nil {
		slog.Error("failed to fetch RSS items",
			slog.String("rss_url", rssURL),
			slog.Any("error", err),
		)
		return Response{}, err
	}

	slog.Info("successfully fetched RSS items",
		slog.Int("item_count", len(titles)),
	)
	slog.Debug("fetched titles", slog.Any("titles", titles))

	return Response{Titles: titles}, nil
}

func parseLogLevel(env string) slog.Level {
	switch env {
	case "DEBUG":
		return slog.LevelDebug
	case "INFO":
		return slog.LevelInfo
	case "WARN":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func main() {
	level := parseLogLevel(os.Getenv("LOG_LEVEL"))

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	}))
	slog.SetDefault(logger)

	slog.Info("logger initialized", slog.String("log_level", level.String()))

	lambda.Start(handler)
}
