package service

import (
	"github.com/gabrielssssssssss/murder-backend.git/internal/model"
	"github.com/gabrielssssssssss/murder-backend.git/internal/repository"
)

type QueryService interface {
	AddIndex(query *model.AddIndexPayload) (*model.AddIndexResponse, error)
	DeleteIndex(query *model.DeleteIndexPayload) (*model.IndexResult, error)
}

type queryServiceimpl struct {
	queryRepository repository.QueryRepository
}

func NewQueryService(queryRepository *repository.QueryRepository) QueryService {
	return &queryServiceimpl{
		queryRepository: *queryRepository,
	}
}
