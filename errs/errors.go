/*
 *  Copyright (c) 2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package errs

import "errors"

var (
	ErrMaxSizeReached     = errors.New("maximum size reached")
	ErrMinSizeInvalid     = errors.New("invalid minimum size")
	ErrPackageSizeInvalid = errors.New("invalid package size")
	ErrHashVerifyFail     = errors.New("fail verify hash")
)
