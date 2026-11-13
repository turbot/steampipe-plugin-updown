package updown

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

type updownConfig struct {
	APIKey *string `hcl:"api_key"`
}

func ConfigInstance() interface{} {
	return &updownConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) updownConfig {
	if connection == nil || connection.GetConfig() == nil {
		return updownConfig{}
	}
	config, _ := connection.GetConfig().(updownConfig)
	return config
}
