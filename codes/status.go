/*
 *  Copyright (c) 2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package codes

type Code uint8

const (
	StatusCodeUnknown Code = 0

	StatusCodePing      Code = 1
	StatusCodeAvailable Code = 2
	StatusCodeReject    Code = 9

	StatusCodeAuthCheck   Code = 20
	StatusCodeAuthSuccess Code = 21
	StatusCodeAuthReject  Code = 22
	StatusCodeAuthFree    Code = 23
	StatusCodeAuthToken   Code = 24

	StatusCodeHandshakeInit    Code = 30
	StatusCodeHandshakeSuccess Code = 31
	StatusCodeHandshakeReject  Code = 32
)
