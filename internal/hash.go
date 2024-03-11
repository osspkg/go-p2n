/*
 *  Copyright (c) 2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package internal

import (
	"crypto/sha256"
)

func Sha256(bb ...[]byte) ([]byte, error) {
	hash := sha256.New()
	for _, b := range bb {
		if _, err := hash.Write(b); err != nil {
			return nil, err
		}
	}
	return hash.Sum(nil), nil
}
