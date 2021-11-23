## v0.1.0 [2021-11-23]

_Enhancements_

- Recompiled plugin with Go version 1.17 ([#8](https://github.com/turbot/steampipe-plugin-updown/pull/8))
- Recompiled plugin with [steampipe-plugin-sdk v1.8.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v182--2021-11-22) ([#7](https://github.com/turbot/steampipe-plugin-updown/pull/7))

## v0.0.3 [2021-09-23]

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk v1.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v161--2021-09-21) ([#3](https://github.com/turbot/steampipe-plugin-updown/pull/3))
- The configuration section in the docs/index.md file has been updated to exclude the need of `api_secret` in the `~/.steampipe/config/updown.spc` file ([#4](https://github.com/turbot/steampipe-plugin-updown/pull/4))

## v0.0.2 [2021-06-16]

_What's new?_

- Return an error if the api_key is not configured.
- Updated plugin license to Apache 2.0 per [turbot/steampipe#488](https://github.com/turbot/steampipe/issues/488)

## v0.0.1 [2021-04-14]

_What's new?_

- New tables added
  - [updown_check](https://hub.steampipe.io/plugins/turbot/updown/tables/updown_check)
  - [updown_downtime](https://hub.steampipe.io/plugins/turbot/updown/tables/updown_downtime)
  - [updown_metric_hour](https://hub.steampipe.io/plugins/turbot/updown/tables/updown_metric_hour)
  - [updown_node](https://hub.steampipe.io/plugins/turbot/updown/tables/updown_node)
