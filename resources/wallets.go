package resources

import (
	"context"

	"github.com/cloudquery/cq-source-plaid/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/plaid/plaid-go/v10/plaid"
)

func Wallets() *schema.Table {
	return &schema.Table{
		Name:      "plaid_wallets",
		Resolver:  fetchWallets,
		Transform: transformers.TransformWithStruct(plaid.Wallet{}, transformers.WithPrimaryKeys("WalletId")),
	}
}

func newWalletListRequest(c *client.Client, cursor string) *plaid.WalletListRequest {
	listRequest := plaid.NewWalletListRequest()
	listRequest.SetClientId(c.ClientId)
	listRequest.SetSecret(c.Secret)
	listRequest.SetCount(20)
	if cursor != "" {
		listRequest.SetCursor(cursor)
	}
	return listRequest
}

func fetchWallets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	client := meta.(*client.Client)
	listRequest := newWalletListRequest(client, "")

	walletsResp, _, err := client.Services.PlaidApi.WalletList(ctx).WalletListRequest(*listRequest).Execute()
	if err != nil {
		return err
	}

	res <- walletsResp.GetWallets()

	if walletsResp.HasNextCursor() {
		listRequest = newWalletListRequest(client, walletsResp.GetNextCursor())
		walletsResp, _, err := client.Services.PlaidApi.WalletList(ctx).WalletListRequest(*listRequest).Execute()
		if err != nil {
			return err
		}
		res <- walletsResp.GetWallets()
	}

	return nil
}
