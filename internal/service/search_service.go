package service

import (
	"github.com/gabrielssssssssss/murder-backend.git/internal/model"
	"github.com/gabrielssssssssss/murder-backend.git/internal/repository"
	"github.com/meilisearch/meilisearch-go"
)

type QueryService interface {
	AddIndex(query *model.AddIndexPayload) (*model.AddIndexResponse, error)
	GetIndex(query *model.IndexPayload) (*meilisearch.Stats, error)
	DeleteIndex(query *model.IndexPayload) (*model.IndexResult, error)
}

type queryServiceimpl struct {
	queryRepository repository.QueryRepository
}

func NewQueryService(queryRepository *repository.QueryRepository) QueryService {
	return &queryServiceimpl{
		queryRepository: *queryRepository,
	}
}
