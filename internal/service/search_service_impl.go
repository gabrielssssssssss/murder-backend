package service

import (
	"fmt"

	"github.com/gabrielssssssssss/murder-backend.git/internal/entity"
	"github.com/gabrielssssssssss/murder-backend.git/internal/model"
	"github.com/meilisearch/meilisearch-go"
)

func (query queryServiceimpl) AddIndex(request *model.AddIndexPayload) (*model.AddIndexResponse, error) {
	if request.Index == "" {
		return nil, fmt.Errorf("Index or document name is empty.")
	}

	input := entity.IndexEntity{
		Name:     request.Index,
		Document: request.Document,
	}

	documentEntity, err := query.queryRepository.AddIndex(&input)
	if err != nil {
		return nil, err
	}

	response := model.AddIndexResponse{
		TaskUid:    documentEntity.TaskUid,
		IndexUid:   documentEntity.IndexUid,
		Status:     documentEntity.Status,
		Type:       documentEntity.Type,
		EnqueuedAt: documentEntity.EnqueuedAt,
	}

	return &response, nil
}

func (query queryServiceimpl) GetIndex(request *model.IndexPayload) (*meilisearch.Stats, error) {
	if request.Index == "" {
		// return nil, fmt.Errorf("Index name is empty.")
	}

	input := entity.IndexEntity{
		Name: request.Index,
	}

	response, err := query.queryRepository.GetIndex(&input)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (query queryServiceimpl) DeleteIndex(request *model.IndexPayload) (*model.IndexResult, error) {
	if request.Index == "" {
		return nil, fmt.Errorf("Index name is empty.")
	}

	input := entity.IndexEntity{
		Name: request.Index,
	}

	response, err := query.queryRepository.DeleteIndex(&input)
	if err != nil {
		return nil, err
	}

	return response, nil
}
