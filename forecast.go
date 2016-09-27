package darksky

type Forecast struct {
	Latitude  float64    `json:"latitude"`
	Longitude float64    `json:"longitude"`
	Timezone  string     `json:"timezone"`
	Currently *DataPoint `json:"currently,omitempty"`
	Minutely  *DataBlock `json:"minutely,omitempty"`
	Hourly    *DataBlock `json:"hourly,omitempty"`
	Daily     *DataBlock `json:"daily,omitempty"`
	Alerts    []Alert    `json:"alerts,omitempty"`
	Flags     *Flags     `json:"flags,omitempty"`
}
