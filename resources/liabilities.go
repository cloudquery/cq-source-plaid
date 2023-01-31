package resources

import (
	"context"

	"github.com/cloudquery/cq-source-plaid/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/plaid/plaid-go/v10/plaid"
)

func Liabilities() *schema.Table {
	return &schema.Table{
		Name:      "plaid_liabilities",
		Resolver:  fetchLiabilities,
		Transform: transformers.TransformWithStruct(plaid.LiabilitiesGetResponse{}),
		Columns: []schema.Column{
			{
				Name:     "item_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Item.ItemId"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchLiabilities(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	client := meta.(*client.Client)

	request := plaid.NewLiabilitiesGetRequest(client.AccessToken)
	resp, _, err := client.Services.PlaidApi.LiabilitiesGet(ctx).LiabilitiesGetRequest(*request).Execute()
	if err != nil {
		return err
	}
	res <- resp

	return nil
}
