---
title: "Steampipe Table: updown_downtime - Query Updown Downtime using SQL"
description: "Allows users to query Updown Downtime, specifically the start and end times, error ratio, and checks affected. It provides insights into service downtime periods and potential issues."
---

# Table: updown_downtime - Query Updown Downtime using SQL

Updown Downtime is a component of the Updown.io monitoring service that records periods of service unavailability or downtime. It provides specific start and end times, the error ratio, and the checks that were affected during these periods. This information can be critical in understanding and addressing service disruptions and maintaining optimal service performance.

## Table Usage Guide

The `updown_downtime` table provides insights into service downtime within Updown.io. As a Site Reliability Engineer, explore downtime-specific details through this table, including start and end times, error ratios, and affected checks. Utilize it to uncover information about service disruptions, such as their duration, the degree of service impact, and the specific checks that were affected.

**Important Notes**
- You must specify the `token` in the `where` clause to query this table.

## Examples

### List all downtimes
Discover the segments that have experienced downtimes in a specific area, ordered chronologically in reverse. This can help identify recurring issues and prioritize areas for improvement.

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