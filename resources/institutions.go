package resources

import (
	"context"

	"github.com/cloudquery/cq-source-plaid/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/plaid/plaid-go/v10/plaid"
)

const count = 500

func Institutions() *schema.Table {
	return &schema.Table{
		Name:      "plaid_institutions",
		Resolver:  fetchInstitutions,
		Transform: transformers.TransformWithStruct(plaid.Institution{}, transformers.WithPrimaryKeys("InstitutionId")),
	}
}

func fetchInstitutions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	client := meta.(*client.Client)

	var total int32
	for {
		request := plaid.NewInstitutionsGetRequest(
			count,
			total,
			[]plaid.CountryCode{
				plaid.COUNTRYCODE_US,
				plaid.COUNTRYCODE_GB,
				plaid.COUNTRYCODE_ES,
				plaid.COUNTRYCODE_NL,
				plaid.COUNTRYCODE_FR,
				plaid.COUNTRYCODE_IE,
				plaid.COUNTRYCODE_CA,
				plaid.COUNTRYCODE_DE,
				plaid.COUNTRYCODE_IT,
				plaid.COUNTRYCODE_PL,
				plaid.COUNTRYCODE_DK,
				plaid.COUNTRYCODE_NO,
				plaid.COUNTRYCODE_SE,
				plaid.COUNTRYCODE_EE,
				plaid.COUNTRYCODE_LT,
				plaid.COUNTRYCODE_LV,
			},
		)
		resp, _, err := client.Services.PlaidApi.InstitutionsGet(ctx).InstitutionsGetRequest(*request).Execute()
		if err != nil {
			return err
		}

		institutions := resp.GetInstitutions()
		total += int32(len(institutions))
		res <- institutions

		if total >= resp.GetTotal() || len(institutions) == 0 {
			break
		}
	}

	return nil
}
