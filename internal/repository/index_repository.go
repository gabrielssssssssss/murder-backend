package repository

import (
	"github.com/gabrielssssssssss/murder-backend.git/internal/entity"
	"github.com/gabrielssssssssss/murder-backend.git/internal/model"
	"github.com/meilisearch/meilisearch-go"
)

type IndexRepository interface {
	NewIndex(query *entity.IndexEntity) (*model.TaskInfo, error)
	AddIndex(query *entity.IndexEntity) (*model.TaskInfo, error)
	GetIndex(query *entity.IndexEntity) (*meilisearch.Stats, error)
	DeleteIndex(query *entity.IndexEntity) (*model.TaskInfo, error)
	GetSettings(query *entity.IndexEntity) (*meilisearch.Settings, error)
	UpdateSettings(query *entity.IndexEntity) (*model.TaskInfo, error)
}

type indexImplementation struct {
	db meilisearch.ServiceManager
}

func NewIndexRepository(client meilisearch.ServiceManager) IndexRepository {
	return &indexImplementation{
		db: client,
	}
}
