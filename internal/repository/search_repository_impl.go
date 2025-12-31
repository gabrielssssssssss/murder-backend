package repository

import (
	"sync"

	"github.com/gabrielssssssssss/murder-backend.git/internal/entity"
	"github.com/meilisearch/meilisearch-go"
)

func (client *searchImplementation) Search(query *entity.SearchEntity) ([]meilisearch.SearchResponse, error) {
	var (
		wg       sync.WaitGroup
		results  = make(chan meilisearch.SearchResponse, len(query.Index))
		response []meilisearch.SearchResponse
	)

	for _, index := range query.Index {
		wg.Add(1)

		go func() {
			defer wg.Done()

			searchRes, err := client.db.Index(index).Search(
				query.Element, &meilisearch.SearchRequest{
					Limit: query.Limit,
				},
			)

			if err != nil {
				return
			}
			results <- *searchRes
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		response = append(response, res)
	}

	return response, nil
}
