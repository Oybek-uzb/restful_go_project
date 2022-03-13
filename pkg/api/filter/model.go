package filter

const (
	DataTypeString = "string"
	DataTypeInt    = "int"
	DataTypeDate   = "date"

	OperatorEqual              = "eq"
	OperatorNotEqual           = "neq"
	OperatorLowerThan          = "lt"
	OperatorLowerThanOrEqual   = "lte"
	OperatorGreaterThan        = "gt"
	OperatorGreaterThanOrEqual = "gte"
	OperatorBetween            = "between"
	OperatorLike               = "like"
)

type options struct {
	Limit  int
	Fields []Field
}

type Field struct {
	Name     string
	Value    string
	Operator string
	Type     string
}

type Options interface {
	GetLimit() int
}
