package repository

import (
	"github.com/gabrielssssssssss/murder-backend.git/internal/entity"
	"github.com/gabrielssssssssss/murder-backend.git/internal/model"
	"github.com/meilisearch/meilisearch-go"
)

type IndexRepository interface {
	AddIndex(query *entity.IndexEntity) (*model.AddIndexResponse, error)
	GetIndex(query *entity.IndexEntity) (*meilisearch.Stats, error)
	DeleteIndex(query *entity.IndexEntity) (*model.IndexResult, error)
}

type indexImplementation struct {
	db meilisearch.ServiceManager
}

func NewIndexRepository(client meilisearch.ServiceManager) IndexRepository {
	return &indexImplementation{
		db: client,
	}
}
