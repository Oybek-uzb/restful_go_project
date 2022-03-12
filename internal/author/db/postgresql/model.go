package postgresql

import (
	"fmt"
	"github.com/Masterminds/squirrel"
)

type sortOptions struct {
	Field, Order string
}

func (so *sortOptions) EnrichQuery(q squirrel.SelectBuilder) squirrel.SelectBuilder {
	return q.OrderBy(fmt.Sprintf("%s %s", so.Field, so.Order))
}
