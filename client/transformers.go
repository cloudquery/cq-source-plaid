package client

import (
	"reflect"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/plaid/plaid-go/v10/plaid"
)

func typeTransformer(field reflect.StructField) (schema.ValueType, error) {
	switch reflect.New(field.Type).Elem().Interface().(type) {
	case plaid.NullableTime:
		return schema.TypeTimestamp, nil
	default:
		return schema.TypeInvalid, nil
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
