package updown

import (
	"context"

	"github.com/antoineaugusti/updown"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableUpDownMetricHour(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "updown_metric_hour",
		Description: "Get detailed metrics about a check. Statistic are aggregated per hour. For example all requests performed between 5:00 and 5:59 will be reported at 5:00 in this API. The time range needs to be at least one hour to get data.",
		List: &plugin.ListConfig{
			Hydrate:    listMetricHour,
			KeyColumns: plugin.SingleColumn("token"),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "token", Type: proto.ColumnType_STRING, Hydrate: tokenString, Transform: transform.FromValue(), Description: "Unique token of the check."},
			{Name: "timestamp", Type: proto.ColumnType_TIMESTAMP, Description: "The time (hour) that the metrics are for."},
			{Name: "apdex", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("Metric.Apdex"), Description: "TODO"},
			{Name: "requests", Type: proto.ColumnType_JSON, Transform: transform.FromField("Metric.Requests"), Description: "Request statistics."},
			{Name: "timings", Type: proto.ColumnType_JSON, Transform: transform.FromField("Metric.Timings"), Description: "Timing statistics."},
			{Name: "host", Type: proto.ColumnType_JSON, Transform: transform.FromField("Metric.Host"), Description: "Host where the check was made."},
		},
	}
}

type metricRow struct {
	Timestamp string
	Metric    updown.MetricItem
}

func listMetricHour(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("updown_metric_hour.listMetricHour", "connection_error", err)
		return nil, err
	}
	quals := d.KeyColumnQuals
	t := quals["token"].GetStringValue()
	result, resp, err := conn.Metric.List(t, "time", "", "")
	if err != nil {
		if resp.StatusCode == 404 {
			// Not found
			return nil, nil
		}
		plugin.Logger(ctx).Error("updown_metric_hour.listMetricHour", "query_error", err, "resp", resp)
		return nil, err
	}
	for ts, i := range result {
		d.StreamListItem(ctx, metricRow{
			Timestamp: ts,
			Metric:    i,
		})
	}
	return nil, nil
}
