# Table: plaid_institutions

The primary key for this table is **institution_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|institution_id (PK)|String|
|name|String|
|products|StringArray|
|country_codes|StringArray|
|url|JSON|
|primary_color|JSON|
|logo|JSON|
|routing_numbers|StringArray|
|oauth|Bool|
|status|JSON|
|payment_initiation_metadata|JSON|
|auth_metadata|JSON|
|additional_properties|JSON|