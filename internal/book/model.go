package book

import "restful_go_project/internal/author"

type Book struct {
	ID     string        `json:"id"`
	Name   string        `json:"name"`
	Author author.Author `json:"author"`
}
