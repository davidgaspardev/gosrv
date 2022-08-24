package model

type Data struct {
	TotalPages uint        `json:"totalPages,omitempty"`
	Data       interface{} `json:"data"`
	Version    string      `json:"version"`
}
