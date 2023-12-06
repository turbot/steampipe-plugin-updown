---
title: "Steampipe Table: updown_metric_hour - Query Updown.io Metric Hours using SQL"
description: "Allows users to query Metric Hours in Updown.io, specifically the hourly checks, responses, and events, providing insights into application uptime and performance."
---

# Table: updown_metric_hour - Query Updown.io Metric Hours using SQL

An Updown.io Metric Hour is a resource that provides an hourly breakdown of checks, responses, and events for a specified Updown.io application. It gives a granular view of application performance and uptime, helping users identify patterns, anomalies, and potential issues. It is part of Updown.io's robust monitoring and reporting capabilities.

## Table Usage Guide

The `updown_metric_hour` table provides detailed insights into the hourly performance and uptime of applications monitored by Updown.io. As a DevOps engineer, you can use this table to explore detailed metrics, including the number of checks, response times, and events that occurred within a specific hour. This can help in identifying trends, diagnosing issues, and ensuring optimal application performance.

**Important Notes**
- You must specify the `token` in the `where` clause to query this table.

## Examples

### List all metrics by hour
Analyze the settings to understand the hourly metrics for a specific token. This can help you track performance trends and identify potential issues swiftly.

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
Determine the instances when the total metric period exceeded 400ms to identify potential performance issues and optimize system efficiency.

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
Analyze the performance of your web services by determining the percentage of response times that are under one second. This can help you assess the speed and efficiency of your services, enabling you to identify potential areas for improvement.

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