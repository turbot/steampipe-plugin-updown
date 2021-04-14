# Table: updown_downtime

History of downtime blocks detected for the checks.

## Examples

### List all downtimes

```sql
select
  *
from
  updown_downtime
order by
  started_at desc
```
