package service

import (
	"fmt"

	"github.com/gabrielssssssssss/murder-backend.git/internal/entity"
	"github.com/gabrielssssssssss/murder-backend.git/internal/model"
	"github.com/meilisearch/meilisearch-go"
)

func (query indexServiceimpl) AddIndex(request *model.AddIndexPayload) (*model.AddIndexResponse, error) {
	if request.Index == "" {
		return nil, fmt.Errorf("Index name is empty.")
	}

	input := entity.IndexEntity{
		Name:       request.Index,
		Delimiter:  request.Delimiter,
		PrimaryKey: request.PrimaryKey,
		Document:   request.Document,
	}

	documentEntity, err := query.indexRepository.AddIndex(&input)
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

func (query indexServiceimpl) GetIndex(request *model.IndexPayload) (*meilisearch.Stats, error) {
	if request.Index == "" {
		return nil, fmt.Errorf("Index name is empty.")
	}

	input := entity.IndexEntity{
		Name: request.Index,
	}

	response, err := query.indexRepository.GetIndex(&input)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (query indexServiceimpl) DeleteIndex(request *model.IndexPayload) (*model.IndexResult, error) {
	if request.Index == "" {
		return nil, fmt.Errorf("Index name is empty.")
	}

	input := entity.IndexEntity{
		Name: request.Index,
	}

	response, err := query.indexRepository.DeleteIndex(&input)
	if err != nil {
		return nil, err
	}

	return response, nil
}
