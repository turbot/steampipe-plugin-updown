package updown

import (
	"context"
	"errors"
	"os"

	"github.com/antoineaugusti/updown"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func connect(_ context.Context, d *plugin.QueryData) (*updown.Client, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "updown"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*updown.Client), nil
	}

	// Default to using env vars
	apiKey := os.Getenv("UPDOWN_API_KEY")

	// But prefer the config
	updownConfig := GetConfig(d.Connection)
	if updownConfig.APIKey != nil {
		apiKey = *updownConfig.APIKey
	}
	
	if apiKey == "" {
		// Credentials not set
		return nil, errors.New("api_key must be configured")
	}

	conn := updown.NewClient(apiKey, nil)

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, conn)

	return conn, nil
}

func tokenString(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	t := quals["token"].GetStringValue()
	return t, nil
}
