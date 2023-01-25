package resources

import (
	"testing"

	_ "embed"

	"github.com/cloudquery/cq-source-plaid/client"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/plaid/plaid-go/v10/plaid"
)

type walletsResponse struct {
	Wallets []plaid.Wallet `json:"wallets"`
}

func TestWallets(t *testing.T) {
	var res walletsResponse
	if err := faker.FakeObject(&res); err != nil {
		t.Fatal(err)
	}
	ts := testServer(t, res)

	defer ts.Close()
	client.TestHelper(t, Wallets(), ts)
}
