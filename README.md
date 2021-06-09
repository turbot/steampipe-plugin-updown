![image](https://hub.steampipe.io/images/plugins/turbot/updown-social-graphic.png)

# updown.io Plugin for Steampipe

Use SQL to query infrastructure including servers, networks, facilities and more from updown.io.

* **[Get started →](https://hub.steampipe.io/plugins/turbot/updown)**
* Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/updown/tables)
* Community: [Discussion forums](https://github.com/turbot/steampipe/discussions)
* Get involved: [Issues](https://github.com/turbot/steampipe-plugin-updown/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):
```shell
steampipe plugin install updown
```

Run a query:
```sql
select
  url,
  uptime,
  down,
  down_since
from
  updown_check
```

## Developing

Prerequisites:
- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone git@github.com:turbot/steampipe-plugin-updown
cd steampipe-plugin-updown
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:
```
make
```

Configure the plugin:
```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/updown.spc
```

Try it!
```
steampipe query
> .inspect updown
```

Further reading:
* [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
* [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-updown/blob/main/LICENSE).

`help wanted` issues:
- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [updown.io Plugin](https://github.com/turbot/steampipe-plugin-updown/labels/help%20wanted)
