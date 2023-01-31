# Table: plaid_investments_holdings

The primary key for this table is **item_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|item_id (PK)|String|
|accounts|JSON|
|holdings|JSON|
|securities|JSON|
|item|JSON|
|request_id|String|
|additional_properties|JSON|