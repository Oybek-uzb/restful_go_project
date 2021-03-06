package model

import (
	"restful_go_project/internal/author/storage"
	"restful_go_project/pkg/api/filter"
)

type filterOptions struct {
	limit  int
	fields []filter.Field
}

func NewFilterOptions(options filter.Options) storage.FilterOptions {
	return &filterOptions{
		limit:  options.Limit(),
		fields: options.Fields(),
	}
}
