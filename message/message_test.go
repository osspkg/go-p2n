/*
 *  Copyright (c) 2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package message_test

import (
	"crypto/cipher"
	"fmt"
	"reflect"
	"testing"

	"go.osspkg.com/p2n/internal"
	"go.osspkg.com/p2n/message"
)

type testModel struct {
	S string
}

func (t *testModel) Marshal() ([]byte, error) {
	return []byte(t.S), nil
}

func (t *testModel) Unmarshal(b []byte) error {
	t.S = string(b)
	return nil
}

func TestUnit_EncodeDecode(t *testing.T) {
	tryTest(t, nil)
}

func TestUnit_EncryptDecrypt(t *testing.T) {
	key, err := internal.BuildKey()
	if err != nil {
		t.Fatal(err)
	}
	tryTest(t, key)
}

func tryTest(t *testing.T, key cipher.AEAD) {
	nonce, err := message.NewNonce()
	if err != nil {
		t.Fatal(err)
	}
	msg := &message.Message{
		Code:    1,
		Chunk:   1,
		Nonce:   nonce,
		Payload: &testModel{S: "Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file."},
	}
	b, err := msg.Encrypt(key)
	if len(b) == 0 || err != nil {
		t.Fatal(err)
	}

	fmt.Print("\n", len(b), "\n", string(b), "\n")

	msg2 := &message.Message{Payload: &testModel{}}
	if err = msg2.Decrypt(b, key); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(msg, msg2) {
		t.Fatal("models not equal")
	}
}
