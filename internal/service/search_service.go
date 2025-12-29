package service

import (
	"github.com/gabrielssssssssss/murder-backend.git/internal/model"
	"github.com/gabrielssssssssss/murder-backend.git/internal/repository"
	"github.com/meilisearch/meilisearch-go"
)

type SearchService interface {
	Search(query *model.SearchPayload) (*meilisearch.SearchResponse, error)
}

type searchServiceimpl struct {
	searchRepository repository.SearchRepository
}

func NewSearchService(queryRepository *repository.SearchRepository) SearchService {
	return &searchServiceimpl{
		searchRepository: *queryRepository,
	}
}
