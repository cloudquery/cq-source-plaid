package plugin

import (
	"github.com/cloudquery/cq-source-plaid/client"
	"github.com/cloudquery/cq-source-plaid/resources"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"plaid",
		Version,
		schema.Tables{
			resources.Transactions(),
			resources.Liabilities(),
			resources.Identities(),
			resources.InvestmentsTransactions(),
			resources.InvestmentsHoldings(),
			resources.AccountBalances(),
			resources.Auths(),
			resources.Wallets(),
			resources.Institutions(),
		},
		client.New,
	)
}
