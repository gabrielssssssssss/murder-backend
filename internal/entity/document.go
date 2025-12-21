package entity

import "mime/multipart"

type IndexEntity struct {
	Name     string
	Document multipart.FileHeader
}
