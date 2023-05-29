package resources

import (
	_ "embed"
	"testing"

	"github.com/cloudquery/cq-source-plaid/client"
	"github.com/cloudquery/plugin-sdk/v3/faker"
	"github.com/plaid/plaid-go/v10/plaid"
)

func TestWallets(t *testing.T) {
	var res plaid.WalletListResponse
	if err := faker.FakeObject(&res); err != nil {
		t.Fatal(err)
	}
	res.SetNextCursor("")
	ts := client.TestServer(t, res)

	defer ts.Close()
	client.TestHelper(t, Wallets(), ts)
}
