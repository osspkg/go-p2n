/*
 *  Copyright (c) 2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package message

import (
	"bytes"
	"crypto/cipher"
	"encoding/binary"

	"go.osspkg.com/p2n/codes"
	"go.osspkg.com/p2n/errs"
	"go.osspkg.com/p2n/internal"
	"go.osspkg.com/p2n/payloads"
)

const (
	MinPackageSize = 1 + 8 + 8 + 32
	MaxPackageSize = (1 << 18) - MinPackageSize
)

type Message struct {
	Code    codes.Code
	Chunk   uint64
	Nonce   Nonce
	Payload payloads.Payload
}

func (v *Message) Encode() (out []byte, err error) {
	return v.Encrypt(nil)
}

func (v *Message) Decode(b []byte) error {
	return v.Decrypt(b, nil)
}

func (v *Message) Encrypt(gcm cipher.AEAD) (out []byte, err error) {
	var pb []byte
	if v.Payload != nil {
		if pb, err = v.Payload.Marshal(); err != nil {
			return nil, err
		}
		if gcm != nil {
			if pb, err = internal.Encrypt(pb, gcm); err != nil {
				return nil, err
			}
		}
		if len(pb) > MaxPackageSize {
			return nil, errs.ErrMaxSizeReached
		}
	}

	var chunk [8]byte
	binary.LittleEndian.PutUint64(chunk[:], v.Chunk)

	hash, err := internal.Sha256(v.Nonce[:], chunk[:], pb)
	if err != nil {
		return nil, err
	}

	out = make([]byte, 0, MinPackageSize+len(pb))
	out = append(out, uint8(v.Code))
	out = append(out, v.Nonce[:]...)
	out = append(out, chunk[:]...)
	out = append(out, hash[:]...)
	out = append(out, pb[:]...)

	return out, nil
}

func (v *Message) Decrypt(b []byte, gcm cipher.AEAD) error {
	if len(b) < MinPackageSize {
		return errs.ErrMinSizeInvalid
	}

	n := 0
	v.Code = codes.Code(b[n])
	n++
	n += copy(v.Nonce[:], b[n:n+8])

	var chunk [8]byte
	n += copy(chunk[:], b[n:n+8])
	v.Chunk = binary.LittleEndian.Uint64(chunk[:])

	var hashM [32]byte
	n += copy(hashM[:], b[n:n+32])

	hashC, err := internal.Sha256(v.Nonce[:], chunk[:], b[n:])
	if err != nil {
		return err
	}

	if !bytes.Equal(hashM[:], hashC) {
		return errs.ErrHashVerifyFail
	}

	if v.Payload == nil {
		return nil
	}

	pb := b[n:]
	if gcm != nil {
		if pb, err = internal.Decrypt(pb, gcm); err != nil {
			return err
		}
	}

	return v.Payload.Unmarshal(pb)
}
