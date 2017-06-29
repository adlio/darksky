Go Dark Sky API
================

[![Powered by Dark Sky](https://raw.githubusercontent.com/adlio/darksky/master/logos/poweredby-oneline.png)](https://darksky.net/poweredby/)

[![GoDoc](https://godoc.org/github.com/adlio/darksky?status.svg)](http://godoc.org/github.com/adlio/darksky)
[![Build Status](https://travis-ci.org/adlio/darksky.svg)](https://travis-ci.org/adlio/darksky)
[![Coverage Status](https://coveralls.io/repos/github/adlio/darksky/badge.svg?branch=master)](https://coveralls.io/github/adlio/darksky?branch=master)

A #golang package to consume [Dark Sky](https://darksky.net) Forecast and Time Machine API calls.

## Getting Started with the Dark Sky API

All usage requires a Dark Sky API key, which you can obtain from the [Dark Sky developer site](https://darksky.net/dev/).

To use the Go libary client, instantiate a `darksky.Client` with your Dark Sky API key:

```Go
lat := "47.202"
lng := "-123.4167"

client := darksky.NewClient("APIKEYHERE")
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
fmt.Printf("Temperature: %.2f\n",f.Currently.Temperature)
fmt.Println("Icon:        " + f.Currently.Icon)
```

## Dark Sky Time Machine Usage

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


## API Arguments

The Dark Sky API accepts a few modification parameters. Set these via a `darksky.Arguments`. If you
want the default behavior, use `darksky.Defaults`. If you're looking only for the `Currently` data
object, then you should use `darksky.CurrentOnly` instead. Examples:

```Go

// No query string parameters (i.e. URL ends with /lat,lng)
f, err := client.GetForecast(lat,lng,darksky.Defaults)

// Currently block only (i.e URL ends with /lat,lng?excludes=minutely,hourly,daily,alerts,flags)
f, err := client.GetForecast(lat,lng,darksky.CurrentOnly)

```

If you'd like to set your own excludes= list, or set other arguments, you'll need to construct a
`darksky.Arguments` directly. The type is a simple map[string]string:

```Go
// Custom query string parameters (/lat,lng?excludes=minutely&units=si&extend=hourly)
f, err := client.GetForecast(lat,lng,darksky.Arguments{"excludes":"minutely","units": "si", "extend": "hourly"})

```

## A Note About time.Time and darksky.Time

The Dark Sky API uses Unix timestamps everywhere times are represented. For #golang developer convenience,
this package uses `time.Time` where possible. The time values inside `Forecast`, however, are of type
`darksky.Time`, which wraps `time.Time`. You can use `.Time` to get the underlying `time.Time` value,
and you can call Time methods directly as well.

```Go
f, err := client.GetTimeMachineForecast(lat, lng, lastYear, darksky.Defaults)

actualTime := f.Currently.Time.Time

fmt.Sprintf("Temp: %.2f\n", f.Currently.Temperature)
fmt.Sprintf("Time: %s\n", f.Currently.Time.Format("2006-01-02 15:04:05"))
fmt.Sprintf("Time: %s\n", actualTime.String())
```

