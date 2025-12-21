package entity

import "mime/multipart"

type DocumentEntity struct {
	Index    string
	Document multipart.FileHeader
}
