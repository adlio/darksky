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
	path := fmt.Sprintf("%s,%s", lat, lng)
	return c.Get(path, args)
}

func (c *Client) GetTimeMachineForecast(lat, lng string, t time.Time, args Arguments) (forecast *Forecast, err error) {
	path := fmt.Sprintf("%s,%s,%d", lat, lng, t.Unix())
	return c.Get(path, args)
}

func (c *Client) Get(path string, args Arguments) (forecast *Forecast, err error) {

	url := fmt.Sprintf("%s/%s/%s", c.BaseURL, c.APIKey, path)

	params := args.ToURLValues()
	if len(params) > 0 {
		url = fmt.Sprintf("%s?%s", url, params.Encode())
	}

	resp, err := c.client.Get(url)
	if err != nil {
		err = errors.Wrapf(err, "HTTP request for /%s request failed.", path)
		return
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&forecast)
	if err != nil {
		err = errors.Wrapf(err, "JSON decode for /%s failed.", path)
		return
	}

	return
}
