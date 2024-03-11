/*
 *  Copyright (c) 2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package internal_test

import (
	"bytes"
	"testing"

	"go.osspkg.com/p2n/internal"
)

func TestNewDH(t *testing.T) {
	dh1, dh2 := internal.NewDH(), internal.NewDH()

	sk1, err := dh1.Generate()
	if err != nil {
		t.Fatal(err)
	}
	sk2, err := dh2.Generate()
	if err != nil {
		t.Fatal(err)
	}

	err1, err2 := dh1.Build(sk2), dh2.Build(sk1)
	if err1 != nil || err2 != nil {
		t.Fatal(err1, err2)
	}

	if !bytes.Equal(dh1.Shared(), dh2.Shared()) {
		t.Fatal("shared keys not equal")
	}
}
