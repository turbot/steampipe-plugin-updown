## v1.1.0 [2025-04-17]

_Dependencies_

- Recompiled plugin with Go version `1.23.1`. ([#50](https://github.com/turbot/steampipe-plugin-updown/pull/50))
- Recompiled plugin with [steampipe-plugin-sdk v5.11.5](https://github.com/turbot/steampipe-plugin-sdk/blob/v5.11.5/CHANGELOG.md#v5115-2025-03-31) that addresses critical and high vulnerabilities in dependent packages. ([#50](https://github.com/turbot/steampipe-plugin-updown/pull/50))

## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#47](https://github.com/turbot/steampipe-plugin-updown/pull/47))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#47](https://github.com/turbot/steampipe-plugin-updown/pull/47))

## v0.6.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#36](https://github.com/turbot/steampipe-plugin-updown/pull/36))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#36](https://github.com/turbot/steampipe-plugin-updown/pull/36))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-updown/blob/main/docs/LICENSE). ([#36](https://github.com/turbot/steampipe-plugin-updown/pull/36))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#35](https://github.com/turbot/steampipe-plugin-updown/pull/35))

## v0.5.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#28](https://github.com/turbot/steampipe-plugin-updown/pull/28))

## v0.5.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#25](https://github.com/turbot/steampipe-plugin-updown/pull/25))
- Recompiled plugin with Go version `1.21`. ([#25](https://github.com/turbot/steampipe-plugin-updown/pull/25))

## v0.4.0 [2023-04-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#18](https://github.com/turbot/steampipe-plugin-updown/pull/18))

## v0.3.0 [2022-09-09]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.6](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v416-2022-09-02) which includes several caching and memory management improvements. ([#16](https://github.com/turbot/steampipe-plugin-updown/pull/16))
- Recompiled plugin with Go version `1.19`. ([#16](https://github.com/turbot/steampipe-plugin-updown/pull/16))

## v0.2.0 [2022-04-28]

_Enhancements_

- Added support for native Linux ARM and Mac M1 builds. ([#12](https://github.com/turbot/steampipe-plugin-updown/pull/12))
- Recompiled plugin with [steampipe-plugin-sdk v3.1.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v310--2022-03-30) and Go version `1.18`. ([#11](https://github.com/turbot/steampipe-plugin-updown/pull/11))

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
