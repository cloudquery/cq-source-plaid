# Table: plaid_transactions

The composite primary key for this table is (**transaction_type**, **transaction_id**, **_transaction_type**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|transaction_type (PK)|String|
|pending_transaction_id|JSON|
|category_id|JSON|
|category|StringArray|
|location|JSON|
|payment_meta|JSON|
|account_owner|JSON|
|name|String|
|original_description|JSON|
|account_id|String|
|amount|Float|
|iso_currency_code|JSON|
|unofficial_currency_code|JSON|
|date|String|
|pending|Bool|
|transaction_id (PK)|String|
|merchant_name|JSON|
|logo_url|JSON|
|website|JSON|
|check_number|JSON|
|payment_channel|String|
|authorized_date|JSON|
|authorized_datetime|Timestamp|
|datetime|Timestamp|
|transaction_code|JSON|
|personal_finance_category|JSON|
|personal_finance_category_icon_url|String|
|counterparties|JSON|
|_transaction_type (PK)|String|