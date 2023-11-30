---
title: "Steampipe Table: updown_node - Query Updown.io Nodes using SQL"
description: "Allows users to query Nodes in Updown.io, specifically the check nodes, providing insights into node status and performance."
---

# Table: updown_node - Query Updown.io Nodes using SQL

Updown.io is a simple and inexpensive website monitoring service that checks your website's status at regular intervals. It provides detailed metrics about your site's uptime, response time, and error analysis. Updown.io allows you to monitor HTTP(s), TCP, and ICMP to ensure your services are always up and running.

## Table Usage Guide

The `updown_node` table provides insights into nodes within Updown.io's monitoring service. As a DevOps engineer, explore node-specific details through this table, including node status, response time, and associated metadata. Utilize it to uncover information about nodes, such as their performance, uptime, and any potential issues or anomalies.

## Examples

### List all nodes
Explore the nodes within your network in a structured order, providing a comprehensive overview to facilitate management and troubleshooting. This query is useful in identifying potential issues or anomalies within your network.

```sql
select
  *
from
  updown_node
order by
  name
```