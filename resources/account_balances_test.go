package resources

import (
	"testing"

	_ "embed"

	"github.com/cloudquery/cq-source-plaid/client"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/plaid/plaid-go/v10/plaid"
)

func TestAccountBalances(t *testing.T) {
	var res plaid.AccountsGetResponse
	if err := faker.FakeObject(&res); err != nil {
		t.Fatal(err)
	}

	ts := client.TestServer(t, res)

	defer ts.Close()
	client.TestHelper(t, AccountBalances(), ts)
}
