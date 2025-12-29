package model

type SearchPayload struct {
	Index   string `form:"index"`
	Element string `form:"element"`
}
