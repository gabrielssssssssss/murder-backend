package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/meilisearch/meilisearch-go"
)

func NewMeilisearchDatabase() meilisearch.ServiceManager {
	MEILISEARCH_HOST := os.Getenv("MEILISEARCH_HOST")
	MEILISEARCH_PORT := os.Getenv("MEILISEARCH_PORT")
	MEILISEARCH_MASTER_KEY := os.Getenv("MEILISEARCH_MASTER_KEY")

	url := fmt.Sprintf("http://%s:%s", MEILISEARCH_HOST, MEILISEARCH_PORT)
	client := meilisearch.New(url, meilisearch.WithAPIKey(MEILISEARCH_MASTER_KEY))

	return client
}

func NewMeilisearchContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 30*time.Second)
}
