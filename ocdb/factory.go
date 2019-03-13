// Copyright 2019 The Alice-Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ocdb

import (
	"reflect"

	"go-hep.org/x/hep/groot/rbytes"
	"go-hep.org/x/hep/groot/root"
	"go-hep.org/x/hep/groot/rtypes"
)

func init() {
	{
		f := func() reflect.Value {
			var o Entry
			return reflect.ValueOf(&o)
		}
		rtypes.Factory.Add("AliCDBEntry", f)
	}
}

var (
	_ root.Object        = (*Entry)(nil)
	_ rbytes.Marshaler   = (*Entry)(nil)
	_ rbytes.Unmarshaler = (*Entry)(nil)
)

func init() {
	{
		f := func() reflect.Value {
			var o ID
			return reflect.ValueOf(&o)
		}
		rtypes.Factory.Add("AliCDBId", f)
	}
}

var (
	_ root.Object        = (*ID)(nil)
	_ rbytes.Marshaler   = (*ID)(nil)
	_ rbytes.Unmarshaler = (*ID)(nil)
)

func init() {
	{
		f := func() reflect.Value {
			var o Path
			return reflect.ValueOf(&o)
		}
		rtypes.Factory.Add("AliCDBPath", f)
	}
}

var (
	_ root.Object        = (*Path)(nil)
	_ rbytes.Marshaler   = (*Path)(nil)
	_ rbytes.Unmarshaler = (*Path)(nil)
)

func init() {
	{
		f := func() reflect.Value {
			var o RunRange
			return reflect.ValueOf(&o)
		}
		rtypes.Factory.Add("AliCDBRunRange", f)
	}
}

var (
	_ root.Object        = (*RunRange)(nil)
	_ rbytes.Marshaler   = (*RunRange)(nil)
	_ rbytes.Unmarshaler = (*RunRange)(nil)
)

func init() {
	{
		f := func() reflect.Value {
			var o MetaData
			return reflect.ValueOf(&o)
		}
		rtypes.Factory.Add("AliCDBMetaData", f)
	}
}

var (
	_ root.Object        = (*MetaData)(nil)
	_ rbytes.Marshaler   = (*MetaData)(nil)
	_ rbytes.Unmarshaler = (*MetaData)(nil)
)
