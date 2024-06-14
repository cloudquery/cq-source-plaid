package resources

import (
	_ "embed"
	"testing"
	"time"

	"github.com/cloudquery/cq-source-plaid/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/plaid/plaid-go/v10/plaid"
)

func TestTransactions(t *testing.T) {
	var res plaid.TransactionsSyncResponse
	if err := faker.FakeObject(&res); err != nil {
		t.Fatal(err)
	}

	res.Added[0].SetPendingTransactionId("pending_transaction_id")
	res.Added[0].SetCategoryId("category_id")
	res.Added[0].SetAccountOwner("account_owner")
	res.Added[0].SetOriginalDescription("original_description")
	res.Added[0].SetIsoCurrencyCode("iso_currency_code")
	res.Added[0].SetUnofficialCurrencyCode("unofficial_currency_code")
	res.Added[0].SetMerchantName("merchant_name")
	res.Added[0].SetLogoUrl("logo_url")
	res.Added[0].SetWebsite("website")
	res.Added[0].SetCheckNumber("check_number")
	res.Added[0].SetAuthorizedDate("authorized_date")
	res.Added[0].SetAuthorizedDatetime(time.Now())
	res.Added[0].SetTransactionCode("transaction_code")
	res.Added[0].SetPersonalFinanceCategory(plaid.PersonalFinanceCategory{})
	res.Added[0].SetDatetime(time.Now())
	res.NextCursor = ""
	res.HasMore = false

	ts := client.TestServer(t, res)

	defer ts.Close()
	client.TestHelper(t, Transactions(), ts)
}
