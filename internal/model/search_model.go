package model

import (
	"mime/multipart"
	"time"
)

type AddIndexPayload struct {
	Index    string               `form:"index"`
	Document multipart.FileHeader `form:"document"`
}

type AddIndexResponse struct {
	TaskUid    int    `json:"taskUid"`
	IndexUid   string `json:"indexUid"`
	Status     string `json:"status"`
	Type       string `json:"type"`
	EnqueuedAt string `json:"enqueuedAt"`
}

type IndexResult struct {
	TaskUID    int64     `json:"taskUid"`
	IndexUID   string    `json:"indexUid"`
	EnqueuedAt time.Time `json:"enqueuedAt"`
}

type DeleteIndexPayload struct {
	Index string `json:"index"`
}
