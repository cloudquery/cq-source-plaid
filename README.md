# CloudQuery Plaid Source Plugin

[![test](https://github.com/cloudquery/cq-source-plaid/actions/workflows/test.yml/badge.svg)](https://github.com/cloudquery/cq-source-plaid/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/cloudquery/cq-source-plaid)](https://goreportcard.com/report/github.com/cloudquery/cq-source-plaid)

A [Plaid](https://plaid.com/) source plugin for CloudQuery that loads data from the [Plaid API](https://plaid.com/docs/api/) to any database, data warehouse or data lake supported by [CloudQuery](https://www.cloudquery.io/docs/quickstart), such as PostgreSQL, BigQuery, Athena, and many more.

## Supported Resources

For a full list of supported resources, see [the tables documentation](./docs/tables/README.md).

## Configuration

The following source configuration file will sync [supported data points](./docs/tables/README.md) to a PostgreSQL database. See [the CloudQuery Quickstart](https://www.cloudquery.io/docs/quickstart) for more information on how to configure the source and destination.

```yaml
kind: source
spec:
  name: "plaid"
  path: "cloudquery/plaid"
  version: "v1.1.0"
  destinations: [postgresql]
  spec:
    # plugin spec section
    client_id: ${PLAID_CLIENT_ID}
    secret: ${PLAID_SECRET}
    access_token: ${PLAID_ACCESS_TOKEN}
    environment: sandbox
```

### Plugin Spec

- `client_id` (string, required):

  A Plaid client ID from your [dashboard](https://dashboard.plaid.com/team/keys). See [the Plaid docs](https://plaid.com/docs/quickstart/#introduction) for more information

- `secret` (string, required):

  A Plaid secret from your [dashboard](https://dashboard.plaid.com/team/keys). See [the Plaid docs](https://plaid.com/docs/quickstart/#introduction) for more information

- `access_token` (string, required):

  A Plaid access token obtained by following the link authorization flow. We provide an [example application](./token-generator/README.md) to generate a token for testing purposes only.
  For production usage you should set up a hosted frontend application and backend server that saves access tokens from link authentication flows initiated by your users.
  See [the Plaid docs](https://plaid.com/docs/link/) for more information

- `environment` (string, optional):

  The Plaid environment to use. Defaults to `sandbox`. See [the Plaid docs](https://plaid.com/docs/api/#api-host) for more information. Should match the Plaid secret you are using

## Example Queries

### List recent transactions

```sql
select name, category, amount, iso_currency_code, date, merchant_name, payment_channel from plaid_transactions
order by
  date desc
limit
  10
```

Example result:

```text
                 name                  |                   category                   | amount | iso_currency_code |    date    |   merchant_name   | payment_channel 
---------------------------------------+----------------------------------------------+--------+-------------------+------------+-------------------+-----------------
 Uber 063015 SF**POOL**                | {Travel,Taxi}                                |    5.4 | "USD"             | 2023-01-30 | "Uber"            | online
 CREDIT CARD 3333 PAYMENT *//          | {Payment,"Credit Card"}                      |     25 | "USD"             | 2023-01-30 | null              | other
 ACH Electronic CreditGUSTO PAY 123456 | {Transfer,Debit}                             |   5850 | "USD"             | 2023-01-29 | null              | online
 CD DEPOSIT .INITIAL.                  | {Transfer,Deposit}                           |   1000 | "USD"             | 2023-01-29 | null              | other
 United Airlines                       | {Travel,"Airlines and Aviation Services"}    |   -500 | "USD"             | 2023-01-28 | "United Airlines" | in store
 Touchstone Climbing                   | {Recreation,"Gyms and Fitness Centers"}      |   78.5 | "USD"             | 2023-01-28 | null              | in store
 Starbucks                             | {"Food and Drink",Restaurants,"Coffee Shop"} |   4.33 | "USD"             | 2023-01-27 | "Starbucks"       | in store
 McDonald's                            | {"Food and Drink",Restaurants,"Fast Food"}   |     12 | "USD"             | 2023-01-27 | "McDonald's"      | in store
 SparkFun                              | {"Food and Drink",Restaurants}               |   89.4 | "USD"             | 2023-01-26 | null              | in store
 INTRST PYMNT                          | {Transfer,Credit}                            |  -4.22 | "USD"             | 2023-01-25 | null              | other
(10 rows)
```

## Development

### Run tests

```bash
make test
```

### Run linter

```bash
make lint
```

### Generate docs

```bash
make gen-docs
```

### Release a new version

1. Follow [this link](https://github.com/cloudquery/cq-source-plaid/releases/new) to draft a new release.
2. Click `Choose a tag` and enter the new version number:
   ![image](https://user-images.githubusercontent.com/26760571/219360662-0ad1f83d-84c9-47c8-afb9-fe774ce03dcc.png)
3. Click `Create new tag: <version> on publish` assuming it's a new tag.
4. Click `Generate release notes` to automatically generate release notes.
5. Click `Publish release` to publish the release.

> Once the tag is pushed, a new GitHub Actions workflow will be triggered to build and upload the release binaries to the release
