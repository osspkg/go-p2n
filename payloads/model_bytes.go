/*
 *  Copyright (c) 2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package payloads

type Bytes struct {
	D []byte
}

func (v *Bytes) Marshal() ([]byte, error) {
	out := make([]byte, len(v.D))
	copy(out, v.D)
	return out, nil
}

func (v *Bytes) Unmarshal(b []byte) error {
	v.D = append(v.D[:0], b...)
	return nil
}
