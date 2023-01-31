package resources

import (
	"testing"

	_ "embed"

	"github.com/cloudquery/cq-source-plaid/client"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/plaid/plaid-go/v10/plaid"
)

func TestInstitutions(t *testing.T) {
	var res plaid.InstitutionsGetResponse
	if err := faker.FakeObject(&res); err != nil {
		t.Fatal(err)
	}
	testString := "test"
	res.Institutions[0].Url.Set(&testString)
	res.Institutions[0].Logo.Set(&testString)
	res.Institutions[0].PrimaryColor.Set(&testString)
	res.Institutions[0].Status.Set(&plaid.InstitutionStatus{})
	res.Institutions[0].PaymentInitiationMetadata.Set(&plaid.PaymentInitiationMetadata{})
	res.Institutions[0].AuthMetadata.Set(&plaid.AuthMetadata{})
	res.SetTotal(1)

	ts := client.TestServer(t, res)

	defer ts.Close()
	client.TestHelper(t, Institutions(), ts)
}
