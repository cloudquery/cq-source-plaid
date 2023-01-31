package resources

import (
	"context"

	"github.com/cloudquery/cq-source-plaid/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/plaid/plaid-go/v10/plaid"
)

func InvestmentsHoldings() *schema.Table {
	return &schema.Table{
		Name:      "plaid_investments_holdings",
		Resolver:  fetchInvestmentsHoldings,
		Transform: transformers.TransformWithStruct(plaid.InvestmentsHoldingsGetResponse{}),
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

func fetchInvestmentsHoldings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	client := meta.(*client.Client)
	request := plaid.NewInvestmentsHoldingsGetRequest(client.AccessToken)
	resp, _, err := client.Services.PlaidApi.InvestmentsHoldingsGet(ctx).InvestmentsHoldingsGetRequest(*request).Execute()
	if err != nil {
		return err
	}
	res <- resp

	return nil
}
