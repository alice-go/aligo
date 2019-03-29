// Copyright 2019 The Alice-Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package muoncalib

import (
	"fmt"
	"reflect"
	"strings"

	"go-hep.org/x/hep/groot/rbase"
	"go-hep.org/x/hep/groot/rbytes"
	"go-hep.org/x/hep/groot/rcont"
	"go-hep.org/x/hep/groot/rdict"
	"go-hep.org/x/hep/groot/rmeta"
	"go-hep.org/x/hep/groot/root"
	"go-hep.org/x/hep/groot/rtypes"
)

type AliMpExMap struct {
	base rbase.Object   `groot:"BASE-TObject"` // base class
	objs rcont.ObjArray `groot:"fObjects"`     // /<  Array of objects
	keys rcont.ArrayL64 `groot:"fKeys"`        // /<  Array of keys
}

func (*AliMpExMap) Class() string   { return "AliMpExMap" }
func (*AliMpExMap) RVersion() int16 { return 1 }

// MarshalROOT implements rbytes.Marshaler
func (o *AliMpExMap) MarshalROOT(w *rbytes.WBuffer) (int, error) {
	if w.Err() != nil {
		return 0, w.Err()
	}

	pos := w.WriteVersion(o.RVersion())

	o.base.MarshalROOT(w)
	o.objs.MarshalROOT(w)
	o.keys.MarshalROOT(w)

	return w.SetByteCount(pos, o.Class())
}

// ROOTUnmarshaler is the interface implemented by an object that can
// unmarshal itself from a ROOT buffer
func (o *AliMpExMap) UnmarshalROOT(r *rbytes.RBuffer) error {
	if r.Err() != nil {
		return r.Err()
	}

	start := r.Pos()
	_, pos, bcnt := r.ReadVersion(o.Class())

	if err := o.base.UnmarshalROOT(r); err != nil {
		return err
	}

	if err := o.objs.UnmarshalROOT(r); err != nil {
		return err
	}

	if err := o.keys.UnmarshalROOT(r); err != nil {
		return err
	}

	r.CheckByteCount(pos, bcnt, start, o.Class())
	return r.Err()
}

func (exmap AliMpExMap) String() string {
	o := new(strings.Builder)
	fmt.Fprintf(o, "ExMap{Objs: [")
	for i := 0; i < exmap.objs.Len(); i++ {
		if i > 0 {
			fmt.Fprintf(o, ", ")
		}
		fmt.Fprintf(o, "%v", exmap.objs.At(i))
	}
	fmt.Fprintf(o, "], Keys: %v}", exmap.keys.Data)
	return o.String()
}

func (e *AliMpExMap) Objects() rcont.ObjArray { return e.objs }
func (e *AliMpExMap) Keys() rcont.ArrayL64    { return e.keys }

func init() {
	{
		f := func() reflect.Value {
			var o AliMpExMap
			return reflect.ValueOf(&o)
		}
		rtypes.Factory.Add("AliMpExMap", f)
	}
}

func init() {
	// Streamer for AliMpExMap.
	rdict.Streamers.Add(rdict.NewCxxStreamerInfo("AliMpExMap", 1, 0x1be35f37, []rbytes.StreamerElement{
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
			Name:   *rbase.NewNamed("fObjects", "/<  Array of objects "),
			Type:   rmeta.Object,
			Size:   64,
			ArrLen: 0,
			ArrDim: 0,
			MaxIdx: [5]int32{0, 0, 0, 0, 0},
			Offset: 0,
			EName:  "TObjArray",
			XMin:   0.000000,
			XMax:   0.000000,
			Factor: 0.000000,
		}.New()},
		&rdict.StreamerObjectAny{StreamerElement: rdict.Element{
			Name:   *rbase.NewNamed("fKeys", "/<  Array of keys "),
			Type:   rmeta.Any,
			Size:   24,
			ArrLen: 0,
			ArrDim: 0,
			MaxIdx: [5]int32{0, 0, 0, 0, 0},
			Offset: 0,
			EName:  "TArrayL",
			XMin:   0.000000,
			XMax:   0.000000,
			Factor: 0.000000,
		}.New()},
	}))
}

var (
	_ root.Object        = (*AliMpExMap)(nil)
	_ rbytes.Marshaler   = (*AliMpExMap)(nil)
	_ rbytes.Unmarshaler = (*AliMpExMap)(nil)
)
