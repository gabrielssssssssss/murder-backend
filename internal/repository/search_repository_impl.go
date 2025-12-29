package repository

import (
	"github.com/gabrielssssssssss/murder-backend.git/internal/entity"
	"github.com/meilisearch/meilisearch-go"
)

func (client *searchImplementation) Search(query *entity.SearchEntity) (*meilisearch.SearchResponse, error) {
	searchRes, err := client.db.Index(query.Index).Search(query.Element,
		&meilisearch.SearchRequest{
			Limit: 20,
		})

	if err != nil {
		return nil, err
	}

	return searchRes, nil
}
