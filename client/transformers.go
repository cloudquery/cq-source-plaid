package client

import (
	"reflect"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/plaid/plaid-go/v10/plaid"
)

func typeTransformer(field reflect.StructField) (arrow.DataType, error) {
	switch reflect.New(field.Type).Elem().Interface().(type) {
	case plaid.NullableTime:
		return arrow.FixedWidthTypes.Timestamp_us, nil
	default:
		return nil, nil
	}
}

func resolverTransformer(field reflect.StructField, path string) schema.ColumnResolver {
	switch reflect.New(field.Type).Elem().Interface().(type) {
	case plaid.NullableTime:
		return ResolveNullableTime(path)
	default:
		return nil
	}
}

func Options() []transformers.StructTransformerOption {
	options := []transformers.StructTransformerOption{
		transformers.WithTypeTransformer(typeTransformer),
		transformers.WithResolverTransformer(resolverTransformer),
	}

	return options
}
