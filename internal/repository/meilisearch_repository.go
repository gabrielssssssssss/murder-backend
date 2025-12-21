package repository

import (
	"github.com/gabrielssssssssss/murder-backend.git/internal/entity"
	"github.com/gabrielssssssssss/murder-backend.git/internal/model"
	"github.com/meilisearch/meilisearch-go"
)

type QueryRepository interface {
	AddIndex(query *entity.IndexEntity) (*model.AddIndexResponse, error)
	DeleteIndex(query *entity.IndexEntity) (*model.IndexResult, error)
}

type queryImplementation struct {
	db meilisearch.ServiceManager
}

func NewQueryRepository(client meilisearch.ServiceManager) QueryRepository {
	return &queryImplementation{
		db: client,
	}
}
