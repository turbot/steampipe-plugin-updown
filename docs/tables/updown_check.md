# Table: updown_check

Query information about checks in the updown.io account.

## Examples

### List all checks

```sql
select
  *
from
  updown_check
```

### Get a check by token

```sql
select
  *
from
  updown_check
where
  token = 'sdfw'
```

### List checks that are currently down

```sql
select
  *
from
  updown_check
where
  down
```
