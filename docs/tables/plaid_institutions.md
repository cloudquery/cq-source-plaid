# Table: plaid_institutions

This table shows data for Plaid Institutions.

The primary key for this table is **institution_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|institution_id (PK)|`utf8`|
|name|`utf8`|
|products|`list<item: utf8, nullable>`|
|country_codes|`list<item: utf8, nullable>`|
|url|`json`|
|primary_color|`json`|
|logo|`json`|
|routing_numbers|`list<item: utf8, nullable>`|
|oauth|`bool`|
|status|`json`|
|payment_initiation_metadata|`json`|
|auth_metadata|`json`|
|additional_properties|`json`|