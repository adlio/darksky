// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package darksky

import (
	"testing"
	"time"
)

func TestGetForecast(t *testing.T) {

	c := testClient()
	c.BaseURL = mockResponse("forecast-response.json").URL
	forecast, err := c.GetForecast("47.202", "-123.4167", Defaults)

	if err != nil {
		t.Fatal(err)
	}

	curr := forecast.Currently

	if curr.Time.UTC().Format("2006-01-02 15:04:05") != "2016-09-28 03:46:47" {
		t.Errorf("Got incorrect time '%s'", curr.Time.UTC())
	}

	if curr.Summary != "Clear" {
		t.Errorf("Expected summary 'Clear'. Got '%s'.", forecast.Currently.Summary)
	}

	if curr.Icon != "clear-night" {
		t.Errorf("Got incorrect icon '%s'.", curr.Icon)
	}

	if curr.NearestStormDistance != 3 {
		t.Errorf("Got incorrect NearestStormDistance %d", curr.NearestStormDistance)
	}

	if curr.Temperature != 55.13 {
		t.Errorf("Got incorrect temperature %.2f.", curr.Temperature)
	}

	minutely := forecast.Minutely
	if len(minutely.Data) != 61 {
		t.Errorf("Expected 61 Minutely entries. Got %d.", len(minutely.Data))
	}
}

func TestGetMinimalForecast(t *testing.T) {
	c := testClient()
	c.BaseURL = mockResponse("minimal-with-alerts.json").URL
	forecast, err := c.GetForecast("38.95663", "-77.396338", CurrentOnly)
	if err != nil {
		t.Fatal(err)
	}

	curr := forecast.Currently

	if curr.Summary != "Clear" {
		t.Errorf("Got incorrect summary '%s'.", curr.Summary)
	}

	alerts := forecast.Alerts
	if len(alerts) != 2 {
		t.Errorf("Expected 2 alerts. Got %d", len(alerts))
	}
	if alerts[0].Title != "Flash Flood Watch for Montgomery, MD" {
		t.Errorf("Incorrect alert title: '%s'.", alerts[0].Title)
	}
}

func TestGetTimeMachine(t *testing.T) {
	c := testClient()
	d := time.Date(2012, 6, 29, 23, 48, 5, 0, time.UTC)
	c.BaseURL = mockResponse("time-machine-response.json").URL
	forecast, err := c.GetTimeMachineForecast("38.95663", "-77.396338", d, Defaults)
	if err != nil {
		t.Error(err)
	}

	curr := forecast.Currently
	if curr.Summary != "Rain and Breezy" {
		t.Errorf("Incorrect summary '%s'.", curr.Summary)
	}

	if curr.Icon != "rain" {
		t.Errorf("Incorrect icon '%s'.", curr.Icon)
	}

	if curr.Temperature != 71.59 {
		t.Errorf("Got incorrect temperature %.2f.", curr.Temperature)
	}

	if forecast.Minutely != nil {
		t.Error("Expected .Minutely to be absent.")
	}
}

func TestGetWithBadURL(t *testing.T) {
	c := testClient()
	c.BaseURL = "gopher://test"
	_, err := c.GetForecast("1.23", "-1.23", Defaults)
	if err == nil {
		t.Fatal("GetForecast() should fail with an invalid URL.")
	}
}

func testClient() *Client {
	c := NewClient("APIKEY")
	return c
}
