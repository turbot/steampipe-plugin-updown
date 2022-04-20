package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-updown/updown"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: updown.Plugin})
}
