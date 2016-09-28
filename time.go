// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package darksky

import (
	"encoding/json"
	"time"
)

type Time struct {
	time.Time
}

func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Unix())
}

func (t *Time) UnmarshalJSON(b []byte) (err error) {
	var src int64
	if err = json.Unmarshal(b, &src); err == nil {
		newTime := time.Unix(src, 0).In(time.UTC)
		*t = Time{newTime}
	}
	return err
}
