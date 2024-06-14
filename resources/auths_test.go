package resources

import (
	_ "embed"
	"testing"

	"github.com/cloudquery/cq-source-plaid/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/plaid/plaid-go/v10/plaid"
)

func TestAuths(t *testing.T) {
	var res plaid.AuthGetResponse
	if err := faker.FakeObject(&res); err != nil {
		t.Fatal(err)
	}

	ts := client.TestServer(t, res)

	defer ts.Close()
	client.TestHelper(t, Auths(), ts)
}
