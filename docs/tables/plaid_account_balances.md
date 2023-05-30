# Table: plaid_account_balances

This table shows data for Plaid Account Balances.

The primary key for this table is **item_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|item_id (PK)|`utf8`|
|accounts|`json`|
|item|`json`|
|request_id|`utf8`|
|additional_properties|`json`|