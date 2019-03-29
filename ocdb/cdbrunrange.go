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

// RunRange represents a [first, last] range of run numbers.
type RunRange struct {
	base  rbase.Object `groot:"BASE-TObject"` // base class
	First int32        `groot:"fFirstRun"`    // first valid run
	Last  int32        `groot:"fLastRun"`     // last valid run
}

func (*RunRange) Class() string   { return "AliCDBRunRange" }
func (*RunRange) RVersion() int16 { return 1 }

// MarshalROOT implements rbytes.Marshaler
func (rr *RunRange) MarshalROOT(w *rbytes.WBuffer) (int, error) {
	if w.Err() != nil {
		return 0, w.Err()
	}

	pos := w.WriteVersion(rr.RVersion())

	rr.base.MarshalROOT(w)
	w.WriteI32(rr.First)
	w.WriteI32(rr.Last)

	return w.SetByteCount(pos, rr.Class())
}

// ROOTUnmarshaler is the interface implemented by an object that can
// unmarshal itself from a ROOT buffer
func (rr *RunRange) UnmarshalROOT(r *rbytes.RBuffer) error {
	if r.Err() != nil {
		return r.Err()
	}

	start := r.Pos()
	_, pos, bcnt := r.ReadVersion(rr.Class())

	if err := rr.base.UnmarshalROOT(r); err != nil {
		return err
	}
	rr.First = r.ReadI32()
	rr.Last = r.ReadI32()

	r.CheckByteCount(pos, bcnt, start, rr.Class())
	return r.Err()
}

func (rr RunRange) String() string {
	return fmt.Sprintf("RunRange{First: %d, Last: %d}", rr.First, rr.Last)
}

func init() {
	// Streamer for AliCDBRunRange.
	rdict.Streamers.Add(rdict.NewCxxStreamerInfo("AliCDBRunRange", 1, 0x8ea5b2d7, []rbytes.StreamerElement{
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
		&rdict.StreamerBasicType{StreamerElement: rdict.Element{
			Name:   *rbase.NewNamed("fFirstRun", "first valid run"),
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
			Name:   *rbase.NewNamed("fLastRun", "last valid run\t"),
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
	}))
}

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
