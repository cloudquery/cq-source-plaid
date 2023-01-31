package resources

import (
	"context"
	"time"

	"github.com/cloudquery/cq-source-plaid/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/plaid/plaid-go/v10/plaid"
)

func InvestmentsTransactions() *schema.Table {
	return &schema.Table{
		Name:      "plaid_investments_transactions",
		Resolver:  fetchInvestmentsTransactions,
		Transform: transformers.TransformWithStruct(plaid.InvestmentsTransactionsGetResponse{}),
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

func fetchInvestmentsTransactions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	client := meta.(*client.Client)
	today := time.Now().Format("2006-01-02")
	request := plaid.NewInvestmentsTransactionsGetRequest(client.AccessToken, "2000-01-01", today)
	resp, _, err := client.Services.PlaidApi.InvestmentsTransactionsGet(ctx).InvestmentsTransactionsGetRequest(*request).Execute()
	if err != nil {
		return err
	}
	res <- resp

	return nil
}
