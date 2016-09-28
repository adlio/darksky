package darksky

import (
	"net/url"
)

type Arguments map[string]string

var Defaults = make(Arguments)
var CurrentOnly = Arguments{"exclude": "minutely,hourly,daily,alerts,flags"}
var CurrentWithAlerts = Arguments{"exclude": "minutely,hourly,daily,flags"}

func (args Arguments) ToURLValues() url.Values {
	v := url.Values{}
	for key, value := range args {
		v.Set(key, value)
	}
	return v
}
