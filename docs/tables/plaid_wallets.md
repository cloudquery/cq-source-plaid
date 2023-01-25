# Table: plaid_wallets

The primary key for this table is **wallet_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|wallet_id (PK)|String|
|balance|JSON|
|numbers|JSON|
|recipient_id|String|
|additional_properties|JSON|