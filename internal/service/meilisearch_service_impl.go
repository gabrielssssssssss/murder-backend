package service

import (
	"fmt"

	"github.com/gabrielssssssssss/murder-backend.git/internal/entity"
	"github.com/gabrielssssssssss/murder-backend.git/internal/model"
)

func (query queryServiceimpl) AddDocument(request *model.AddDocumentQueryPayload) (*model.AddDocumentQueryResponse, error) {
	if request.Index == "" {
		return nil, fmt.Errorf("Index or document name is empty.")
	}

	input := entity.DocumentEntity{
		Index:    request.Index,
		Document: request.Document,
	}

	documentEntity, err := query.queryRepository.AddDocument(&input)
	if err != nil {
		return nil, err
	}

	response := model.AddDocumentQueryResponse{
		TaskUid:    documentEntity.TaskUid,
		IndexUid:   documentEntity.IndexUid,
		Status:     documentEntity.Status,
		Type:       documentEntity.Type,
		EnqueuedAt: documentEntity.EnqueuedAt,
	}

	return &response, nil
}

func (query queryServiceimpl) DelDocument(request *model.DelDocumentQueryPayload) (bool, error) {
	if request.Index == "" {
		return false, fmt.Errorf("Index name is empty.")
	}

	input := entity.DocumentEntity{
		Index: request.Index,
	}

	_, err := query.queryRepository.DelDocument(&input)
	if err != nil {
		return false, err
	}

	return true, nil
}
