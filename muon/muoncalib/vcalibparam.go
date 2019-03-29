// Copyright 2019 The Alice-Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package muoncalib

import (
	"reflect"

	"go-hep.org/x/hep/groot/rbase"
	"go-hep.org/x/hep/groot/rbytes"
	"go-hep.org/x/hep/groot/rdict"
	"go-hep.org/x/hep/groot/rmeta"
	"go-hep.org/x/hep/groot/root"
	"go-hep.org/x/hep/groot/rtypes"
)

type AliMUONVCalibParam struct {
	base rbase.Object `groot:"BASE-TObject"` // base class
}

func (*AliMUONVCalibParam) Class() string   { return "AliMUONVCalibParam" }
func (*AliMUONVCalibParam) RVersion() int16 { return 1 }

// MarshalROOT implements rbytes.Marshaler
func (o *AliMUONVCalibParam) MarshalROOT(w *rbytes.WBuffer) (int, error) {
	if w.Err() != nil {
		return 0, w.Err()
	}

	pos := w.WriteVersion(o.RVersion())

	o.base.MarshalROOT(w)

	return w.SetByteCount(pos, o.Class())
}

// ROOTUnmarshaler is the interface implemented by an object that can
// unmarshal itself from a ROOT buffer
func (o *AliMUONVCalibParam) UnmarshalROOT(r *rbytes.RBuffer) error {
	if r.Err() != nil {
		return r.Err()
	}

	start := r.Pos()
	_, pos, bcnt := r.ReadVersion(o.Class())

	if err := o.base.UnmarshalROOT(r); err != nil {
		return err
	}

	r.CheckByteCount(pos, bcnt, start, o.Class())
	return r.Err()
}

func init() {
	{
		f := func() reflect.Value {
			var o AliMUONVCalibParam
			return reflect.ValueOf(&o)
		}
		rtypes.Factory.Add("AliMUONVCalibParam", f)
	}
}

func init() {
	// Streamer for AliMUONVCalibParam.
	rdict.Streamers.Add(rdict.NewCxxStreamerInfo("AliMUONVCalibParam", 0, 0xfbc20a45, []rbytes.StreamerElement{
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
	}))
}

var (
	_ root.Object        = (*AliMUONVCalibParam)(nil)
	_ rbytes.Marshaler   = (*AliMUONVCalibParam)(nil)
	_ rbytes.Unmarshaler = (*AliMUONVCalibParam)(nil)
)
