package service

import (
	"github.com/gabrielssssssssss/murder-backend.git/internal/model"
	"github.com/gabrielssssssssss/murder-backend.git/internal/repository"
)

type QueryService interface {
	AddDocument(query *model.AddDocumentQueryPayload) (*model.AddDocumentQueryResponse, error)
	DelDocument(query *model.DelDocumentQueryPayload) (bool, error)
}

type queryServiceimpl struct {
	queryRepository repository.QueryRepository
}

func NewQueryService(queryRepository *repository.QueryRepository) QueryService {
	return &queryServiceimpl{
		queryRepository: *queryRepository,
	}
}
