// Code generated by goa v3.3.1, DO NOT EDIT.
//
// weather client
//
// Command:
// $ goa gen github.com/penkovski/goademo/design

package weather

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "weather" service client.
type Client struct {
	WeatherQueryEndpoint goa.Endpoint
}

// NewClient initializes a "weather" service client given the endpoints.
func NewClient(weatherQuery goa.Endpoint) *Client {
	return &Client{
		WeatherQueryEndpoint: weatherQuery,
	}
}

// WeatherQuery calls the "WeatherQuery" endpoint of the "weather" service.
func (c *Client) WeatherQuery(ctx context.Context, p *WeatherQueryRequest) (res *WeatherQueryResult, err error) {
	var ires interface{}
	ires, err = c.WeatherQueryEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*WeatherQueryResult), nil
}
