/*
 *  Copyright (c) 2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package message

import "crypto/rand"

type Nonce [8]byte

func NewNonce() (Nonce, error) {
	var v Nonce
	_, err := rand.Read(v[:])
	return v, err
}
