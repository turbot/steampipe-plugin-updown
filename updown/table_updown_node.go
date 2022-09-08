package updown

import (
	"context"

	"github.com/antoineaugusti/updown"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableUpDownNode(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "updown_node",
		Description: "List all updown.io monitoring nodes.",
		List: &plugin.ListConfig{
			Hydrate: listNode,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "IPv4 address of the node."},
			{Name: "ip", Type: proto.ColumnType_IPADDR, Transform: transform.FromField("Node.IP"), Description: "IPv4 address of the node."},
			{Name: "ip6", Type: proto.ColumnType_IPADDR, Transform: transform.FromField("Node.IP6"), Description: "IPv6 address of the node."},
			{Name: "city", Type: proto.ColumnType_STRING, Transform: transform.FromField("Node.City"), Description: "City of the node."},
			{Name: "country", Type: proto.ColumnType_STRING, Transform: transform.FromField("Node.Country"), Description: "Country of the node."},
			{Name: "country_code", Type: proto.ColumnType_STRING, Transform: transform.FromField("Node.CountryCode"), Description: "Country code of the node."},
		},
	}
}

type nodeRow struct {
	Name string
	Node updown.NodeDetails
}

func listNode(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("updown_node.listNode", "connection_error", err)
		return nil, err
	}
	result, resp, err := conn.Node.List()
	if err != nil {
		plugin.Logger(ctx).Error("updown_node.listNode", "query_error", err, "resp", resp)
		return nil, err
	}
	for name, i := range result {
		d.StreamListItem(ctx, nodeRow{
			Name: name,
			Node: i,
		})
	}
	return nil, nil
}
