// Copyright (c) 2013 wathiede[0]. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
//	[0]: https://github.com/wathiede

package fileutil

import (
	"os"
)

// PunchHole deallocates space inside a file in the byte range starting at
// offset and continuing for len bytes. Unimplemented on FreeBSD.
func PunchHole(f *os.File, off, len int64) error {
	return nil
}

// Fadvise predeclares an access pattern for file data.  See also 'man 2
// posix_fadvise'. Unimplemented on FreeBSD.
func Fadvise(f *os.File, off, len int64, advice FadviseAdvice) error {
	return nil
}
