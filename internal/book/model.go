package book

import (
	"restful_go_project/internal/author/model"
)

type Book struct {
	ID     string       `json:"id"`
	Name   string       `json:"name"`
	Author model.Author `json:"author"`
}
