package entity

import "mime/multipart"

type IndexEntity struct {
	Name       string
	Delimiter  string
	PrimaryKey string
	Document   multipart.FileHeader
}
