package updown

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-updown",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		DefaultTransform: transform.FromGo().NullIfZero(),
		TableMap: map[string]*plugin.Table{
			"updown_check":       tableUpDownCheck(ctx),
			"updown_downtime":    tableUpDownDowntime(ctx),
			"updown_metric_hour": tableUpDownMetricHour(ctx),
			"updown_node":        tableUpDownNode(ctx),
		},
	}
	return p
}
