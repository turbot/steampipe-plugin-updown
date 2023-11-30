---
title: "Steampipe Table: updown_check - Query Updown.io Checks using SQL"
description: "Allows users to query Checks in Updown.io, specifically the status, uptime, and downtime of each check, providing insights into website performance and availability."
---

# Table: updown_check - Query Updown.io Checks using SQL

Updown.io is a website monitoring service that checks your website's status every few seconds from different locations around the world. It provides detailed reports of uptime, downtime, and performance metrics. This service helps you stay informed about the health and performance of your website and take appropriate actions when predefined conditions are met.

## Table Usage Guide

The `updown_check` table provides insights into Checks within Updown.io. As a Site Reliability Engineer, explore check-specific details through this table, including uptime, downtime, and associated metadata. Utilize it to uncover information about checks, such as those with significant downtime, the performance metrics of each check, and the verification of uptime reports.

## Examples

### List all checks
Explore all monitoring checks in your system to better understand the health and performance of your digital assets. This can help you proactively identify potential issues and rectify them before they escalate.

```sql
select
  *
from
  updown_check
```

### Get a check by token
Discover the segments that correspond to a specific identifier in order to analyze and manage system checks more efficiently. This can be useful in pinpointing specific checks for troubleshooting or system optimization.

```sql
select
  *
from
  updown_check
where
  token = 'sdfw'
```

### List checks that are currently down
Identify instances where certain checks are currently down. This can assist in promptly addressing and rectifying any issues that might be affecting system performance or functionality.

```sql
select
  *
from
  updown_check
where
  down
```