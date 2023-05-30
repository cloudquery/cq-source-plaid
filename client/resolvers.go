package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/plaid/plaid-go/v10/plaid"
	"github.com/thoas/go-funk"
)

func ResolveNullableTime(path string) schema.ColumnResolver {
	return func(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
		data := funk.Get(resource.Item, path)
		if data == nil {
			return nil
		}
		ts, ok := data.(plaid.NullableTime)
		if !ok {
			return fmt.Errorf("unexpected type, wanted \"*plaid.NullableTime\", have \"%T\"", data)
		}
		return resource.Set(c.Name, ts.Get())
	}
}
