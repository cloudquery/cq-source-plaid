# Table: plaid_transactions

This table shows data for Plaid Transactions.

The composite primary key for this table is (**transaction_type**, **transaction_id**, **_transaction_type**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|transaction_type (PK)|`utf8`|
|pending_transaction_id|`json`|
|category_id|`json`|
|category|`list<item: utf8, nullable>`|
|location|`json`|
|payment_meta|`json`|
|account_owner|`json`|
|name|`utf8`|
|original_description|`json`|
|account_id|`utf8`|
|amount|`float64`|
|iso_currency_code|`json`|
|unofficial_currency_code|`json`|
|date|`utf8`|
|pending|`bool`|
|transaction_id (PK)|`utf8`|
|merchant_name|`json`|
|logo_url|`json`|
|website|`json`|
|check_number|`json`|
|payment_channel|`utf8`|
|authorized_date|`json`|
|authorized_datetime|`timestamp[us, tz=UTC]`|
|datetime|`timestamp[us, tz=UTC]`|
|transaction_code|`json`|
|personal_finance_category|`json`|
|personal_finance_category_icon_url|`utf8`|
|counterparties|`json`|
|_transaction_type (PK)|`utf8`|