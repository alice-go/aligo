// Copyright 2019 The Alice-Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ocdb

import "fmt"

func (rr RunRange) String() string {
	return fmt.Sprintf("RunRange{First: %d, Last: %d}", rr.First, rr.Last)
}
