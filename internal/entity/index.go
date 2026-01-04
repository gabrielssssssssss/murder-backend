package entity

import (
	"mime/multipart"

	"github.com/meilisearch/meilisearch-go"
)

type IndexEntity struct {
	Name       string
	Delimiter  string
	PrimaryKey string
	Document   multipart.FileHeader
	Settings   meilisearch.Settings
}
