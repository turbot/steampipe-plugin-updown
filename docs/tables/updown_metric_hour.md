# Table: updown_metric_hour

Hourly metrics collected for the given check. Results are limited to the last month only.

Note: The `token` field must be set in the `where` clause.

## Examples

### List all metrics by hour

```sql
select
  *
from
  updown_metric_hour
where
  token = '3sdv'
order by
  timestamp desc
```

### List all metric periods where the total time was greater than 400ms

```sql
select
  timestamp,
  timings ->> 'total' as timing_total
from
  updown_metric_hour
where
  token = '3sdv'
  and (timings ->> 'total')::int > 400
order by
  timestamp desc
```

### Percentage of samples responding under 1 second

```sql
select
  100 * (requests -> 'by_response_time' -> 'under1000')::int / (requests -> 'samples')::float as req_under_1sec
from
  updown_metric_hour
where
  token = '3sdv'
order by
  timestamp desc
```
