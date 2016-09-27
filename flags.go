package darksky

type Flags struct {
	DarkSkyUnavailable bool     `json:"darksky-unavailable"`
	MetNoLicense       bool     `json:"meto-license"`
	Sources            []string `json:"sources"`
}
