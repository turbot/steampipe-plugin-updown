package updown

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type updownConfig struct {
	APIKey *string `hcl:"api_key"`
}

func ConfigInstance() interface{} {
	return &updownConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) updownConfig {
	if connection == nil || connection.Config == nil {
		return updownConfig{}
	}
	config, _ := connection.Config.(updownConfig)
	return config
}
