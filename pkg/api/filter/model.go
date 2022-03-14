package filter

import "fmt"

const (
	DataTypeString = "string"
	DataTypeInt    = "int"
	DataTypeDate   = "date"
	DataTypeBool   = "bool"

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
	limit  int
	fields []Field
}

func NewOptions(limit int) Options {
	return &options{limit: limit}
}

type Field struct {
	Name     string
	Value    string
	Operator string
	Type     string
}

type Options interface {
	Limit() int
	AddField(name, operator, value, dataType string) error
	Fields() []Field
}

func (o *options) Limit() int {
	return o.limit
}

func (o *options) AddField(name, operator, value, dataType string) error {
	err := validateOperator(operator)
	if err != nil {
		return err
	}
	o.fields = append(o.fields, Field{
		Name:     name,
		Value:    value,
		Operator: operator,
		Type:     dataType,
	})
	return nil
}

func (o *options) Fields() []Field {
	return o.fields
}

func validateOperator(operator string) error {
	switch operator {
	case OperatorEqual,
		OperatorNotEqual,
		OperatorLowerThan,
		OperatorLowerThanOrEqual,
		OperatorGreaterThan,
		OperatorGreaterThanOrEqual:
		return nil
	default:
		return fmt.Errorf("wrong operator")
	}
}
