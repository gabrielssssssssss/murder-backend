package repository

import "github.com/meilisearch/meilisearch-go"

type SearchRepository interface {
}

type searchImplementation struct {
	db meilisearch.ServiceManager
}

// func NewSearchRepository(client meilisearch.ServiceManager) {
// 	return &searchImplementation{
// 		db: client,
// 	}
// }
