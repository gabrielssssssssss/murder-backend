package repository

import (
	"github.com/gabrielssssssssss/murder-backend.git/internal/entity"
	"github.com/gabrielssssssssss/murder-backend.git/internal/model"
	"github.com/meilisearch/meilisearch-go"
)

type QueryRepository interface {
	AddDocument(query *entity.DocumentEntity) (*model.AddDocumentQueryResponse, error)
	DelDocument(query *entity.DocumentEntity) (bool, error)
}

type queryImplementation struct {
	db meilisearch.ServiceManager
}

func NewQueryRepository(client meilisearch.ServiceManager) QueryRepository {
	return &queryImplementation{
		db: client,
	}
}
