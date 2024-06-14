# Table: plaid_investments_transactions

This table shows data for Plaid Investments Transactions.

The primary key for this table is **item_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|item_id (PK)|`utf8`|
|item|`json`|
|accounts|`json`|
|securities|`json`|
|investment_transactions|`json`|
|total_investment_transactions|`int64`|
|request_id|`utf8`|
|additional_properties|`json`|