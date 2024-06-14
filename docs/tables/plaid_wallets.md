# Table: plaid_wallets

This table shows data for Plaid Wallets.

The primary key for this table is **wallet_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|wallet_id (PK)|`utf8`|
|balance|`json`|
|numbers|`json`|
|recipient_id|`utf8`|
|additional_properties|`json`|