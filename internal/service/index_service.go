package service

import (
	"github.com/gabrielssssssssss/murder-backend.git/internal/model"
	"github.com/gabrielssssssssss/murder-backend.git/internal/repository"
	"github.com/meilisearch/meilisearch-go"
)

type IndexService interface {
	NewIndex(query *model.IndexPayload) (*model.TaskInfo, error)
	AddIndex(query *model.AddIndexPayload) (*model.TaskInfo, error)
	GetIndex(query *model.IndexPayload) (*meilisearch.Stats, error)
	DeleteIndex(query *model.IndexPayload) (*model.IndexResult, error)
	GetSettings(query *model.IndexPayload) (*meilisearch.Settings, error)
}

type indexServiceimpl struct {
	indexRepository repository.IndexRepository
}

func NewIndexService(queryRepository *repository.IndexRepository) IndexService {
	return &indexServiceimpl{
		indexRepository: *queryRepository,
	}
}
