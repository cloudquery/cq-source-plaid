package resources

import (
	"context"

	"github.com/cloudquery/cq-source-plaid/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/plaid/plaid-go/v10/plaid"
)

func newTransactionsSyncRequest(c *client.Client, cursor string) *plaid.TransactionsSyncRequest {
	listRequest := plaid.NewTransactionsSyncRequest(c.AccessToken)
	listRequest.SetCount(500)
	if cursor != "" {
		listRequest.SetCursor(cursor)
	}
	return listRequest
}

func Transactions() *schema.Table {
	return &schema.Table{
		Name:     "plaid_transactions",
		Resolver: fetchTransactions,
		Transform: transformers.TransformWithStruct(transaction{}, append(
			client.Options(),
			transformers.WithPrimaryKeys("TransactionType", "TransactionId"),
			transformers.WithUnwrapStructFields("Transaction"))...,
		),
	}
}

type transaction struct {
	plaid.Transaction
	TransactionType string `json:"_transaction_type"`
}

func saveTransactions(resp plaid.TransactionsSyncResponse, res chan<- any) {
	for _, t := range resp.Added {
		res <- transaction{
			Transaction:     t,
			TransactionType: "added",
		}
	}

	for _, t := range resp.Modified {
		res <- transaction{
			Transaction:     t,
			TransactionType: "modified",
		}
	}

	for _, t := range resp.Removed {
		res <- transaction{
			Transaction:     plaid.Transaction{TransactionId: *t.TransactionId},
			TransactionType: "removed",
		}
	}
}

func fetchTransactions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	request := newTransactionsSyncRequest(cl, "")
	resp, _, err := cl.Services.PlaidApi.TransactionsSync(ctx).TransactionsSyncRequest(*request).Execute()
	if err != nil {
		return err
	}
	saveTransactions(resp, res)

	for resp.HasMore {
		request = newTransactionsSyncRequest(cl, resp.GetNextCursor())
		resp, _, err = cl.Services.PlaidApi.TransactionsSync(ctx).TransactionsSyncRequest(*request).Execute()
		if err != nil {
			return err
		}
		saveTransactions(resp, res)
	}

	return nil
}
