---
organization: Turbot
category: ["internet"]
icon_url: "/images/plugins/turbot/updown.svg"
brand_color: "#55CC44"
display_name: updown.io
name: updown
description: Steampipe plugin for querying updown.io checks, metrics and downtime data.
og_description: Query updown.io with SQL! Open source CLI. No DB required.
og_image: "/images/plugins/turbot/updown-social-graphic.png"
---

# updown.io + Steampipe

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

[updown](https://updown.io) is an online service that checks your website's status by periodically sending an HTTP request to the URL of your choice. It then notifies you by email or sms when your website is not responding correctly.

For example:
```sql
select
  url,
  uptime,
  down,
  down_since
from
  updown_check
```

```
+----------------------+--------+-------+------------+
| url                  | uptime | down  | down_since |
+----------------------+--------+-------+------------+
| https://steampipe.io | 100    | false | <null>     |
| https://turbot.com   | 100    | false | <null>     |
+----------------------+--------+-------+------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/updown/tables)**

## Get started

### Install

Download and install the latest updown.io plugin:

```bash
steampipe plugin install updown
```

### Credentials

| Item | Description |
| - | - |
| Credentials | [Create an API Key](https://updown.io/settings/edit). |
| Permissions | Use a `Read-only` key. |
| Radius | Each connection represents a single updown.io account. |
| Resolution |  1. `api_key` in Steampipe config.<br />2. `UPDOWN_API_KEY` environment variable. |

### Configuration

Installing the latest updown plugin will create a config file (`~/.steampipe/config/updown.spc`) with a single connection named `updown`:

```hcl
connection "updown" {
  plugin     = "updown"
  api_key    = "9m_kAcfuTlW_JCrvoMYK6g"
}
```

## Get involved

* Open source: https://github.com/turbot/steampipe-plugin-updown
* Community: [Discussion forums](https://github.com/turbot/steampipe/discussions)
