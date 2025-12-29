package model

type SearchPayload struct {
	Index   []string `json:"index"`
	Element string   `json:"element"`
}
