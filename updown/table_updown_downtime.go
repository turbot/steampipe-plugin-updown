package updown

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableUpDownDowntime(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "updown_downtime",
		Description: "Downtime represents a downtime period for a check.",
		List: &plugin.ListConfig{
			Hydrate:    listDowntime,
			KeyColumns: plugin.SingleColumn("token"),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "token", Type: proto.ColumnType_STRING, Hydrate: tokenString, Transform: transform.FromValue(), Description: "Unique token of the check."},
			{Name: "error", Type: proto.ColumnType_STRING, Description: "Error from the downtime, if any."},
			{Name: "started_at", Type: proto.ColumnType_TIMESTAMP, Description: "Time when the downtime started."},
			{Name: "ended_at", Type: proto.ColumnType_TIMESTAMP, Description: "Time when the downtime ended."},
			{Name: "duration", Type: proto.ColumnType_INT, Description: "Duration of the downtime."},
		},
	}
}

func listDowntime(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("updown_downtime.listDowntime", "connection_error", err)
		return nil, err
	}
	quals := d.EqualsQuals
	t := quals["token"].GetStringValue()
	// Rough paging approach - best we can do given the API
	page := 1
	for page < 100 {
		items, resp, err := conn.Downtime.List(t, 1)
		if err != nil {
			if resp.StatusCode == 404 {
				// Not found
				return nil, nil
			}
			plugin.Logger(ctx).Error("updown_downtime.listDowntime", "query_error", err, "resp", resp)
			return nil, err
		}
		if len(items) == 0 {
			// No more items, so abort paging. This is silly, but there is
			// no other way to tell.
			break
		}
		for _, i := range items {
			d.StreamListItem(ctx, i)
		}
		page++
	}
	return nil, nil
}
