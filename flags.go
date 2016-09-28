// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package darksky

type Flags struct {
	DarkSkyUnavailable bool     `json:"darksky-unavailable"`
	MetNoLicense       bool     `json:"meto-license"`
	Sources            []string `json:"sources"`
}
