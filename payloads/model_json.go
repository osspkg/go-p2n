/*
 *  Copyright (c) 2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package payloads

import "encoding/json"

type Json struct {
	D interface{}
}

func (v *Json) Marshal() ([]byte, error) {
	return json.Marshal(v.D)
}

func (v *Json) Unmarshal(b []byte) error {
	return json.Unmarshal(b, v.D)
}
