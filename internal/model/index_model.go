package model

import (
	"mime/multipart"
)

type AddIndexPayload struct {
	Index      string               `form:"index"`
	Delimiter  string               `form:"delimiter"`
	PrimaryKey string               `form:"primaryKey"`
	Document   multipart.FileHeader `form:"document"`
}

type TaskInfo struct {
	TaskUid    int    `json:"taskUid"`
	IndexUid   string `json:"indexUid"`
	Status     string `json:"status"`
	Type       string `json:"type"`
	EnqueuedAt string `json:"enqueuedAt"`
}

type IndexPayload struct {
	Index string `json:"index"`
}
