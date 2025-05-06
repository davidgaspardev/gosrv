package model

type Result struct {
	TotalPages uint   `json:"totalPages,omitempty"`
	Result     any    `json:"result"`
	Version    string `json:"version"`
}
