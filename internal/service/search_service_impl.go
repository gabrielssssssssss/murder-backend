package service

import (
	"fmt"

	"github.com/gabrielssssssssss/murder-backend.git/internal/entity"
	"github.com/gabrielssssssssss/murder-backend.git/internal/model"
	"github.com/meilisearch/meilisearch-go"
)

func (query *searchServiceimpl) Search(request *model.SearchPayload) (*meilisearch.MultiSearchResponse, error) {
	if request.Element == "" || request.Index == nil {
		return nil, fmt.Errorf("Index or element is empty.")
	}

	input := entity.SearchEntity{
		Index:   request.Index,
		Element: request.Element,
	}

	response, err := query.searchRepository.Search(&input)
	if err != nil {
		return nil, err
	}

	return response, nil
}
