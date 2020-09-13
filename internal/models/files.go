package models

type Files struct {
	Id   uint   `json:"id"`
	Type string `json:"type"`
	Url  string `json:"url"`
	Size uint64 `json:"size"`
}
