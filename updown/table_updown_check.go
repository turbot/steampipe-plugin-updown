package updown

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableUpDownCheck(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "updown_check",
		Description: "Checks configured for the account.",
		List: &plugin.ListConfig{
			Hydrate: listCheck,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getCheck,
			KeyColumns: plugin.SingleColumn("token"),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "token", Type: proto.ColumnType_STRING, Description: "The check unique token."},
			{Name: "url", Type: proto.ColumnType_STRING, Description: "The URL being monitored."},
			{Name: "alias", Type: proto.ColumnType_STRING, Description: "Alias for the check."},
			{Name: "last_status", Type: proto.ColumnType_INT, Description: "Last status of the check."},
			{Name: "uptime", Type: proto.ColumnType_DOUBLE, Description: "Uptime (percentage) for the check."},
			{Name: "down", Type: proto.ColumnType_BOOL, Description: "True if the URL is down."},
			{Name: "down_since", Type: proto.ColumnType_TIMESTAMP, Description: "The URL has been down since this time. Null if the URL is up."},
			{Name: "error", Type: proto.ColumnType_STRING, Description: "Error from the check, if any."},
			{Name: "period", Type: proto.ColumnType_INT, Description: "Interval in seconds (15, 30, 60, 120, 300, 600, 1800 or 3600)."},
			{Name: "apdex_threshold", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("Apdex"), Description: "APDEX (Application Performance Index) threshold in seconds (0.125, 0.25, 0.5, 1.0 or 2.0)."},
			{Name: "enabled", Type: proto.ColumnType_BOOL, Description: "True if the check is enabled."},
			{Name: "published", Type: proto.ColumnType_BOOL, Description: "True if the status of the check is public."},
			{Name: "last_check_at", Type: proto.ColumnType_TIMESTAMP, Description: "Time of the last check."},
			{Name: "next_check_at", Type: proto.ColumnType_TIMESTAMP, Description: "Time when the next check will be performed."},
			{Name: "favicon_url", Type: proto.ColumnType_STRING, Transform: transform.FromField("FaviconURL"), Description: "Favicon for the URL."},
			{Name: "ssl", Type: proto.ColumnType_JSON, Transform: transform.FromField("SSL"), Description: "SSL information."},
			{Name: "mute_until", Type: proto.ColumnType_TIMESTAMP, Description: "Check is muted until this time."},
			{Name: "disabled_locations", Type: proto.ColumnType_JSON, Description: "Disabled monitoring locations. It's an array of abbreviated location names. Can be any of these: lan, mia, bhs, rbx, fra, sin, tok, syd."},
		},
	}
}

func listCheck(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("updown_check.listCheck", "connection_error", err)
		return nil, err
	}
	result, resp, err := conn.Check.List()
	if err != nil {
		plugin.Logger(ctx).Error("updown_check.listCheck", "query_error", err, "resp", resp)
		return nil, err
	}
	for _, i := range result {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getCheck(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("updown_check.getCheck", "connection_error", err)
		return nil, err
	}
	quals := d.KeyColumnQuals
	token := quals["token"].GetStringValue()
	result, resp, err := conn.Check.Get(token)
	if err != nil {
		if resp.StatusCode == 404 {
			// Not found
			return nil, nil
		}
		plugin.Logger(ctx).Error("updown_check.getCheck", "query_error", err, "resp", resp)
		return nil, err
	}
	return result, nil
}
