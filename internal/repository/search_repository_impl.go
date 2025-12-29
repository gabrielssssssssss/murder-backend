package repository

import (
	"github.com/gabrielssssssssss/murder-backend.git/internal/entity"
	"github.com/meilisearch/meilisearch-go"
)

func (client *searchImplementation) Search(query *entity.SearchEntity) (*meilisearch.MultiSearchResponse, error) {
	queriesPayload := []*meilisearch.SearchRequest{}

	for _, index := range query.Index {
		queriesPayload = append(queriesPayload, &meilisearch.SearchRequest{
			IndexUID: index,
			Query:    query.Element,
			Limit:    10,
		})
	}

	searchRes, err := client.db.MultiSearch(&meilisearch.MultiSearchRequest{
		Queries: queriesPayload,
	})

	if err != nil {
		return nil, err
	}

	return searchRes, nil
}
