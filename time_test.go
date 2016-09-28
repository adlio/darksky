// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package darksky

import (
	"testing"
	"time"
)

func TestMarshalJSON(t *testing.T) {
	t1 := Time{time.Date(2016, 2, 24, 1, 2, 3, 0, time.UTC)}
	str1, err := t1.MarshalJSON()
	if err != nil {
		t.Error(err)
	}

	if string(str1) != "1456275723" {
		t.Errorf("Incorrect time format '%s'.", str1)
	}
}

func TestUnmarshalJSON(t *testing.T) {
	t1 := Time{}
	err := t1.UnmarshalJSON([]byte("1456275723"))
	if err != nil {
		t.Error(err)
	}

	t2 := Time{time.Date(2016, 2, 24, 1, 2, 3, 0, time.UTC)}
	if t1 != t2 {
		t.Errorf("UnmarshalJSON issue '%s' vs '%s'.", t1, t2)
	}

}
