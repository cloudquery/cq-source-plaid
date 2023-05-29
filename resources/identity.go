package resources

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cq-source-plaid/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/plaid/plaid-go/v10/plaid"
)

func Identities() *schema.Table {
	return &schema.Table{
		Name:      "plaid_identities",
		Resolver:  fetchIdentities,
		Transform: transformers.TransformWithStruct(plaid.IdentityGetResponse{}),
		Columns: []schema.Column{
			{
				Name:       "item_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Item.ItemId"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchIdentities(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	client := meta.(*client.Client)

	request := plaid.NewIdentityGetRequest(client.AccessToken)
	resp, _, err := client.Services.PlaidApi.IdentityGet(ctx).IdentityGetRequest(*request).Execute()
	if err != nil {
		return err
	}
	res <- resp

	return nil
}
