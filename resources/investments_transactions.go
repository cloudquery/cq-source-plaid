package resources

import (
	"context"
	"time"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/cloudquery/cq-source-plaid/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/plaid/plaid-go/v10/plaid"
)

func InvestmentsTransactions() *schema.Table {
	return &schema.Table{
		Name:      "plaid_investments_transactions",
		Resolver:  fetchInvestmentsTransactions,
		Transform: transformers.TransformWithStruct(plaid.InvestmentsTransactionsGetResponse{}),
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

func fetchInvestmentsTransactions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	today := time.Now().Format("2006-01-02")
	request := plaid.NewInvestmentsTransactionsGetRequest(cl.AccessToken, "2000-01-01", today)
	resp, _, err := cl.Services.PlaidApi.InvestmentsTransactionsGet(ctx).InvestmentsTransactionsGetRequest(*request).Execute()
	if err != nil {
		return err
	}
	res <- resp

	return nil
}
