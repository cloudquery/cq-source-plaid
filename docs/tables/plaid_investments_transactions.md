# Table: plaid_investments_transactions

The primary key for this table is **item_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|item_id (PK)|String|
|item|JSON|
|accounts|JSON|
|securities|JSON|
|investment_transactions|JSON|
|total_investment_transactions|Int|
|request_id|String|
|additional_properties|JSON|