package model

import "mime/multipart"

type AddDocumentQueryPayload struct {
	Index    string               `form:"index"`
	Document multipart.FileHeader `form:"document"`
}

type DelDocumentQueryPayload struct {
	Index string `json:"index"`
}

type AddDocumentQueryResponse struct {
	TaskUid    int    `json:"taskUid"`
	IndexUid   string `json:"indexUid"`
	Status     string `json:"status"`
	Type       string `json:"type"`
	EnqueuedAt string `json:"enqueuedAt"`
}

type DeleteDocumentQueryResponse struct{}
