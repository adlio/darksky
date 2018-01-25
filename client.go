// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package darksky

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"context"

	"github.com/pkg/errors"
)

type Client struct {
	client  *http.Client
	BaseURL string
	APIKey  string
}

const DEFAULT_BASEURL = "https://api.darksky.net/forecast"

func NewClient(apiKey string) *Client {
	return &Client{
		client:  http.DefaultClient,
		APIKey:  apiKey,
		BaseURL: DEFAULT_BASEURL,
	}
}

func (c *Client) GetForecast(lat, lng string, args Arguments) (forecast *Forecast, err error) {
	return c.GetForecastCtx(context.Background(), lat, lng, args)
}

func (c *Client) GetForecastCtx(ctx context.Context, lat, lng string, args Arguments) (forecast *Forecast, err error) {
	path := fmt.Sprintf("%s,%s", lat, lng)
	return c.GetCtx(ctx, path, args)
}

func (c *Client) GetTimeMachineForecast(lat, lng string, t time.Time, args Arguments) (forecast *Forecast, err error) {
	return c.GetTimeMachineForecastCtx(context.Background(), lat, lng, t, args)
}

func (c *Client) GetTimeMachineForecastCtx(ctx context.Context, lat, lng string, t time.Time, args Arguments) (forecast *Forecast, err error) {
	path := fmt.Sprintf("%s,%s,%d", lat, lng, t.Unix())
	return c.GetCtx(ctx, path, args)
}

func (c *Client) Get(path string, args Arguments) (Forecast *Forecast, err error) {
	return c.GetCtx(context.Background(), path, args)
}

func (c *Client) GetCtx(ctx context.Context, path string, args Arguments) (forecast *Forecast, err error) {

	url := fmt.Sprintf("%s/%s/%s", c.BaseURL, c.APIKey, path)

	params := args.ToURLValues()
	if len(params) > 0 {
		url = fmt.Sprintf("%s?%s", url, params.Encode())
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		err = errors.Wrap(err, "create request")
		return
	}

	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		err = errors.Wrapf(err, "HTTP request for /%s request failed.", path)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = errors.Errorf("darksky API responded with %s", resp.Status)
		return
	}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&forecast)
	if err != nil {
		err = errors.Wrapf(err, "JSON decode for /%s failed.", path)
		return
	}

	return
}
