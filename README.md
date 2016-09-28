Go Dark Sky API
================

[![Powered by Dark Sky](https://raw.githubusercontent.com/adlio/darksky/master/logos/poweredby-oneline.png)](https://darksky.net/poweredby/)

[![GoDoc](https://godoc.org/github.com/adlio/darksky?status.svg)](http://godoc.org/github.com/adlio/darksky)
[![Build Status](https://travis-ci.org/adlio/darksky.svg)](https://travis-ci.org/adlio/darksky)
[![Coverage Status](https://coveralls.io/repos/github/adlio/darksky/badge.svg?branch=master)](https://coveralls.io/github/adlio/darksky?branch=master)

A #golang package to consume [Dark Sky](https://darksky.net) Forecast and Time Machine API calls.

## Example Usage

All usage requires a Dark Sky API key, which you can obtain from the [Dark Sky developer site](https://darksky.net/dev/).

To use the Go libary client, instantiate a `darksky.Client` with your Dark Sky API key:

```Go
lat := "47.202"
lng := "-123.4167"

client := darksky.Client("APIKEYHERE")
f, err := client.GetForecast(lat, lng, darksky.Defaults)
if err != nil {
  // Handle error
}
```

See the [Forecast](https://github.com/adlio/darksky/blob/master/forecast.go),
[DataBlock](https://github.com/adlio/darksky/blob/master/data-block.go), and
[DataPoint](https://github.com/adlio/darksky/blob/master/data-point.go) structs to
get a picture of the shape of the returned data.

You may also want to explore the [Dark Sky Response Format](https://darksky.net/dev/docs/response)
documentation, which explains when each property is expected to be populated (note that DataPoint)
will be very sparse for certain kinds of output.

Chances are, you're looking for the current temperature and a weather summary. Get those thusly:

```Go
fmt.Println("Summary:     " + f.Currently.Summary)
fmt.Println("Temperature: " + f.Currently.Temperature)
fmt.Println("Icon:        " + f.Currently.Icon)
```

## Time Machine Usage

The [Dark Sky](https://darksky.net) API supports requests for retrieving weather data
in the past and the future through time machine calls. Use `GetTimeMachineForecast(lat, lng, time, args)`
to retrieve time machine data.

```Go
lat := "47.202"
lng := "-123.4167"
lastYear := time.Now().AddDate(-1,0,0)

client := darksky.Client("APIKEYHERE")
f, err := client.GetTimeMachineForecast(lat, lng, lastYear, darksky.Defaults)
if err != nil {
  // Handle error
}
```

## A Note About time.Time and darksky.Time

The Dark Sky API uses Unix timestamps everywhere times are represented. For #golang developer convenience,
this package uses `time.Time` where possible. The time values inside `Forecast`, however, are of type
`darksky.Time`, which wraps `time.Time`. You can use `.Time()` to get the underlying `time.Time` value,
and you can call Time methods directly as well.

```Go
f, err := client.GetTimeMachineForecast(lat, lng, lastYear, darksky.Defaults)
fmt.Println("Temp: " + f.Currently.Temperature)
fmt.Println("Time: " + f.Currently.Time.Format("2006-01-02 15:04:05"))
```

