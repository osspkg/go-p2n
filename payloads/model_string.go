/*
 *  Copyright (c) 2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package payloads

type Strings struct {
	D string
}

func (v *Strings) Marshal() ([]byte, error) {
	return []byte(v.D), nil
}

func (v *Strings) Unmarshal(b []byte) error {
	v.D = string(b)
	return nil
}
