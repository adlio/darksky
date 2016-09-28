// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package darksky

type Alert struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URI         string `json:"uri"`
	Expires     int    `json:"expires"`
}
