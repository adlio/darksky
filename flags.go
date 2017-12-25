// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package darksky

type Flags struct {
	DarkSkyUnavailable string   `json:"darksky-unavailable"`
	Sources            []string `json:"sources"`
	Units              string   `json:"units"`
}
