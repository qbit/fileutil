// Copyright (c) 2013 jnml. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build: go1.3

package fileutil

import (
	"io"
	"os"
)

// PunchHole deallocates space inside a file in the byte range starting at
// offset and continuing for len bytes. Not supported on Solaris.
func PunchHole(f *os.File, off, len int64) error {
	return nil
}

// Fadvise predeclares an access pattern for file data.  See also 'man 2
// posix_fadvise'. Not supported on Solaris.
func Fadvise(f *os.File, off, len int64, advice FadviseAdvice) error {
	return nil
}

// IsEOF reports whether err is an EOF condition.
func IsEOF(err error) bool { return err == io.EOF }
