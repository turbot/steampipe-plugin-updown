# Table: updown_downtime

History of downtime blocks detected for the checks.

Note: The `token` field must be set in the `where` clause.

## Examples

### List all downtimes

```sql
select
  *
from
  updown_downtime
where
  token = '3sdv'
order by
  started_at desc
```
