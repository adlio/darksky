// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package darksky

import (
	"testing"
)

func TestToURLValues(t *testing.T) {
	a := Arguments{"exclude": "minutely"}
	values := a.ToURLValues()

	if values.Encode() != "exclude=minutely" {
		t.Errorf("ToURLValues() encoded incorrectly: '%s'.", values.Encode())
	}
}
