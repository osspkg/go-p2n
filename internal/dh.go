/*
 *  Copyright (c) 2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package internal

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"math/big"
)

type DH struct {
	key    *ecdsa.PrivateKey
	shared []byte
}

func NewDH() *DH {
	return &DH{}
}

func (v *DH) Generate() (sub []byte, err error) {
	v.key, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return v.Generate()
	}
	x, y := v.key.X.Bytes(), v.key.Y.Bytes()
	sub = append(x, y...)
	if len(sub) != key64 {
		return v.Generate()
	}
	return
}

func (v *DH) Build(b []byte) error {
	if len(b) != key64 {
		return fmt.Errorf("invalid sub key len")
	}
	x, y := (new(big.Int)).SetBytes(b[:key32]), (new(big.Int)).SetBytes(b[key32:])
	x1, y1 := v.key.PublicKey.ScalarMult(x, y, v.key.D.Bytes())
	s1, s2 := x1.Bytes(), y1.Bytes()

	switch true {
	case len(s1) == key32:
		v.shared = s1
	case len(s2) == key32:
		v.shared = s2
	default:
		return fmt.Errorf("invalid shared key len")
	}

	return nil
}

func (v *DH) Shared() []byte {
	return v.shared
}
