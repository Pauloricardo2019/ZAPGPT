package model

type Choice struct {
	Index   int     `json:"index"`
	Message Message `json:"message"`
}
