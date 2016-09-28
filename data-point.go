// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package darksky

type DataPoint struct {
	ApparentTemperature        float64 `json:"apparentTemperature,omitempty"`
	ApparentTemperatureMax     float64 `json:"apparentTemperatureMax,omitempty"`
	ApparentTemperatureMaxTime Time    `json:"apparentTemperatureMaxTime,omitempty"`
	ApparentTemperatureMin     float64 `json:"apparentTemperatureMin,omitempty"`
	ApparentTemperatureMinTime Time    `json:"apparentTemperatureMinTime,omitempty"`
	CloudCover                 float64 `json:"cloudCover,omitempty"`
	DewPoint                   float64 `json:"dewPoint,omitempty"`
	Humidity                   float64 `json:"humidity,omitempty"`
	Icon                       string  `json:"icon,omitempty"`
	MoonPhase                  float64 `json:"moonPhase,omitempty"`
	NearestStormBearing        float64 `json:"nearestStormBearing,omitempty"`
	NearestStormDistance       float64 `json:"nearestStormDistance,omitempty"`
	Ozone                      float64 `json:"ozone,omitempty"`
	PrecipAccumulation         float64 `json:"precipAccumulation"`
	PrecipIntensity            float64 `json:"precipIntensity"`
	PrecipIntensityMax         float64 `json:"precipIntensityMax"`
	PrecipIntensityMaxTime     Time    `json:"precipIntensityMaxTime"`
	PrecipProbability          float64 `json:"precipProbability"`
	PrecipType                 string  `json:"precipType"`
	Pressure                   float64 `json:"pressure"`
	Summary                    string  `json:"summary"`
	SunriseTime                float64 `json:"sunriseTime"`
	SunsetTime                 float64 `json:"sunsetTime"`
	Temperature                float64 `json:"temperature"`
	TemperatureMax             float64 `json:"temperatureMax"`
	TemperatureMaxTime         float64 `json:"temperatureMaxTime"`
	TemperatureMin             float64 `json:"temperatureMin"`
	TemperatureMinTime         float64 `json:"temperatureMinTime"`
	Time                       Time    `json:"time"`
	Visibility                 float64 `json:"visibility"`
	WindBearing                float64 `json:"windBearing"`
	WindSpeed                  float64 `json:"windSpeed"`
}
