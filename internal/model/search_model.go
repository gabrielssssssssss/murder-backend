package model

type SearchPayload struct {
	Index   []string `json:"index"`
	Limit   int64    `json:"limit"`
	Element string   `json:"element"`
}
