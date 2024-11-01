package config

const StoragePath = "./sqlite/storage.db"

type Motivation struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

var ErrorResponse struct {
	Error string `json:"error,omitempty"`
}
