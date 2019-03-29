// Copyright 2019 The Alice-Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ocdb

import (
	"fmt"
	"reflect"

	"go-hep.org/x/hep/groot/rbase"
	"go-hep.org/x/hep/groot/rbytes"
	"go-hep.org/x/hep/groot/rdict"
	"go-hep.org/x/hep/groot/rmeta"
	"go-hep.org/x/hep/groot/root"
	"go-hep.org/x/hep/groot/rtypes"
)

// ID uniquely identifies an entry in an OCDB file.
type ID struct {
	base    rbase.Object `groot:"BASE-TObject"` // base class
	path    Path         `groot:"fPath"`        // path
	runs    RunRange     `groot:"fRunRange"`    // run range
	vers    int32        `groot:"fVersion"`     // version
	subvers int32        `groot:"fSubVersion"`  // subversion
	last    string       `groot:"fLastStorage"` // previous storage place (new, grid, local, dump)
}

func (*ID) Class() string   { return "AliCDBId" }
func (*ID) RVersion() int16 { return 1 }

// MarshalROOT implements rbytes.Marshaler
func (id *ID) MarshalROOT(w *rbytes.WBuffer) (int, error) {
	if w.Err() != nil {
		return 0, w.Err()
	}

	pos := w.WriteVersion(id.RVersion())

	id.base.MarshalROOT(w)
	id.path.MarshalROOT(w)
	id.runs.MarshalROOT(w)
	w.WriteI32(id.vers)
	w.WriteI32(id.subvers)
	w.WriteString(id.last)

	return w.SetByteCount(pos, id.Class())
}

// ROOTUnmarshaler is the interface implemented by an object that can
// unmarshal itself from a ROOT buffer
func (id *ID) UnmarshalROOT(r *rbytes.RBuffer) error {
	if r.Err() != nil {
		return r.Err()
	}

	start := r.Pos()
	_, pos, bcnt := r.ReadVersion(id.Class())

	if err := id.base.UnmarshalROOT(r); err != nil {
		return err
	}
	if err := id.path.UnmarshalROOT(r); err != nil {
		return err
	}
	if err := id.runs.UnmarshalROOT(r); err != nil {
		return err
	}
	id.vers = r.ReadI32()
	id.subvers = r.ReadI32()
	id.last = r.ReadString()

	r.CheckByteCount(pos, bcnt, start, id.Class())
	return r.Err()
}

func (id ID) Runs() RunRange {
	return id.runs
}

func (id ID) String() string {
	return fmt.Sprintf("AliCDBId{Path: %v, RunRange: %v, Version: 0x%x, SubVersion: 0x%x, Last: %q}",
		id.path, id.runs, id.vers, id.subvers, id.last,
	)
}

func init() {
	{
		f := func() reflect.Value {
			var o ID
			return reflect.ValueOf(&o)
		}
		rtypes.Factory.Add("AliCDBId", f)
	}
}

func init() {
	// Streamer for AliCDBId.
	rdict.Streamers.Add(rdict.NewCxxStreamerInfo("AliCDBId", 1, 0x3308a8aa, []rbytes.StreamerElement{
		rdict.NewStreamerBase(rdict.Element{
			Name:   *rbase.NewNamed("TObject", "Basic ROOT object"),
			Type:   rmeta.Base,
			Size:   0,
			ArrLen: 0,
			ArrDim: 0,
			MaxIdx: [5]int32{0, -1877229523, 0, 0, 0},
			Offset: 0,
			EName:  "BASE",
			XMin:   0.000000,
			XMax:   0.000000,
			Factor: 0.000000,
		}.New(), 1),
		&rdict.StreamerObject{StreamerElement: rdict.Element{
			Name:   *rbase.NewNamed("fPath", "path\t"),
			Type:   rmeta.Object,
			Size:   120,
			ArrLen: 0,
			ArrDim: 0,
			MaxIdx: [5]int32{0, 0, 0, 0, 0},
			Offset: 0,
			EName:  "AliCDBPath",
			XMin:   0.000000,
			XMax:   0.000000,
			Factor: 0.000000,
		}.New()},
		&rdict.StreamerObject{StreamerElement: rdict.Element{
			Name:   *rbase.NewNamed("fRunRange", "run range"),
			Type:   rmeta.Object,
			Size:   24,
			ArrLen: 0,
			ArrDim: 0,
			MaxIdx: [5]int32{0, 0, 0, 0, 0},
			Offset: 0,
			EName:  "AliCDBRunRange",
			XMin:   0.000000,
			XMax:   0.000000,
			Factor: 0.000000,
		}.New()},
		&rdict.StreamerBasicType{StreamerElement: rdict.Element{
			Name:   *rbase.NewNamed("fVersion", "version"),
			Type:   rmeta.Int,
			Size:   4,
			ArrLen: 0,
			ArrDim: 0,
			MaxIdx: [5]int32{0, 0, 0, 0, 0},
			Offset: 0,
			EName:  "Int_t",
			XMin:   0.000000,
			XMax:   0.000000,
			Factor: 0.000000,
		}.New()},
		&rdict.StreamerBasicType{StreamerElement: rdict.Element{
			Name:   *rbase.NewNamed("fSubVersion", "subversion"),
			Type:   rmeta.Int,
			Size:   4,
			ArrLen: 0,
			ArrDim: 0,
			MaxIdx: [5]int32{0, 0, 0, 0, 0},
			Offset: 0,
			EName:  "Int_t",
			XMin:   0.000000,
			XMax:   0.000000,
			Factor: 0.000000,
		}.New()},
		&rdict.StreamerString{StreamerElement: rdict.Element{
			Name:   *rbase.NewNamed("fLastStorage", "previous storage place (new, grid, local, dump)"),
			Type:   rmeta.TString,
			Size:   24,
			ArrLen: 0,
			ArrDim: 0,
			MaxIdx: [5]int32{0, 0, 0, 0, 0},
			Offset: 0,
			EName:  "TString",
			XMin:   0.000000,
			XMax:   0.000000,
			Factor: 0.000000,
		}.New()},
	}))
}

var (
	_ root.Object        = (*ID)(nil)
	_ rbytes.Marshaler   = (*ID)(nil)
	_ rbytes.Unmarshaler = (*ID)(nil)
)
