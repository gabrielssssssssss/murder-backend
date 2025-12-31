package repository

import (
	"github.com/gabrielssssssssss/murder-backend.git/internal/entity"
	"github.com/meilisearch/meilisearch-go"
)

type SearchRepository interface {
	Search(query *entity.SearchEntity) ([]meilisearch.SearchResponse, error)
}

type searchImplementation struct {
	db meilisearch.ServiceManager
}

func NewSearchRepository(client meilisearch.ServiceManager) SearchRepository {
	return &searchImplementation{
		db: client,
	}
}
