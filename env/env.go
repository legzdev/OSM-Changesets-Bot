package env

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

const (
	DataDir = "data"
)

var (
	BotToken  string
	ChannelID int64

	FeedURL     string
	DataBaseURL string

	ChecksInterval  time.Duration
	RetriesInterval time.Duration
)

type ErrMissingEnv struct {
	env string
}

func (err *ErrMissingEnv) Error() string {
	return fmt.Sprintf("missing or invalid env: %s", err.env)
}

func Load() (err error) {
	godotenv.Load()

	err = os.MkdirAll(DataDir, os.ModePerm)
	if err != nil {
		return err
	}

	BotToken = os.Getenv("BOT_TOKEN")
	if BotToken == "" {
		return &ErrMissingEnv{env: "BOT_TOKEN"}
	}

	channelID := os.Getenv("CHANNEL_ID")

	ChannelID, err = strconv.ParseInt(channelID, 10, 64)
	if err != nil {
		return &ErrMissingEnv{env: "CHANNEL_ID"}
	}

	DataBaseURL = os.Getenv("DATABASE_URL")
	if DataBaseURL == "" {
		DataBaseURL = path.Join(DataDir, "database.db")
	}

	FeedURL = os.Getenv("FEED_URL")
	if FeedURL == "" {
		return &ErrMissingEnv{env: "FEED_URL"}
	}

	taskInterval := os.Getenv("CHECKS_INTERVAL")
	if taskInterval == "" {
		ChecksInterval = 1 * time.Minute
	} else {
		ChecksInterval, err = time.ParseDuration(taskInterval)
		if err != nil {
			return &ErrMissingEnv{env: "CHECKS_INTERVAL"}
		}
	}

	retriesInterval := os.Getenv("RETRIES_INTERVAL")
	if retriesInterval == "" {
		RetriesInterval = 5 * time.Second
	} else {
		RetriesInterval, err = time.ParseDuration(retriesInterval)
		if err != nil {
			return &ErrMissingEnv{env: "RETRIES_INTERVAL"}
		}
	}

	return nil
}
