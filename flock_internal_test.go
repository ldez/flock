// Copyright 2015 Tim Heckman. All rights reserved.
// Copyright 2018-2024 The Gofrs. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package flock

import (
	"os"
	"testing"
)

func Test(t *testing.T) {
	tmpFileFh, _ := os.CreateTemp(os.TempDir(), "go-flock-")
	tmpFileFh.Close()
	tmpFile := tmpFileFh.Name()
	os.Remove(tmpFile)

	lock := New(tmpFile)
	locked, err := lock.TryLock()
	if locked == false || err != nil {
		t.Fatalf("failed to lock: locked: %t, err: %v", locked, err)
	}

	newLock := New(tmpFile)
	locked, err = newLock.TryLock()
	if locked != false || err != nil {
		t.Fatalf("should have failed locking: locked: %t, err: %v", locked, err)
	}

	if newLock.fh != nil {
		t.Fatal("file handle should have been released and be nil")
	}
}
