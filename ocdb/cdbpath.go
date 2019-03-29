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

// Path represents a provenance path in an OCDB file.
type Path struct {
	base     rbase.Object `groot:"BASE-TObject"` // base class
	path     string       `groot:"fPath"`        // detector pathname (Detector/DBType/SpecType)
	lvl0     string       `groot:"fLevel0"`      // level0 name (ex. detector: ZDC, TPC...)
	lvl1     string       `groot:"fLevel1"`      // level1 name (ex. DB type, Calib, Align)
	lvl2     string       `groot:"fLevel2"`      // level2 name (ex. DetSpecType, pedestals, gain...)
	valid    bool         `groot:"fIsValid"`     // validity flag
	wildcard bool         `groot:"fIsWildCard"`  // wildcard flag
}

func (*Path) Class() string   { return "AliCDBPath" }
func (*Path) RVersion() int16 { return 1 }

// MarshalROOT implements rbytes.Marshaler
func (p *Path) MarshalROOT(w *rbytes.WBuffer) (int, error) {
	if w.Err() != nil {
		return 0, w.Err()
	}

	pos := w.WriteVersion(p.RVersion())

	p.base.MarshalROOT(w)
	w.WriteString(p.path)
	w.WriteString(p.lvl0)
	w.WriteString(p.lvl1)
	w.WriteString(p.lvl2)
	w.WriteBool(p.valid)
	w.WriteBool(p.wildcard)

	return w.SetByteCount(pos, p.Class())
}

// ROOTUnmarshaler is the interface implemented by an object that can
// unmarshal itself from a ROOT buffer
func (p *Path) UnmarshalROOT(r *rbytes.RBuffer) error {
	if r.Err() != nil {
		return r.Err()
	}

	start := r.Pos()
	_, pos, bcnt := r.ReadVersion(p.Class())

	if err := p.base.UnmarshalROOT(r); err != nil {
		return err
	}

	p.path = r.ReadString()
	p.lvl0 = r.ReadString()
	p.lvl1 = r.ReadString()
	p.lvl2 = r.ReadString()
	p.valid = r.ReadBool()
	p.wildcard = r.ReadBool()

	r.CheckByteCount(pos, bcnt, start, p.Class())
	return r.Err()
}

func (p Path) String() string {
	return fmt.Sprintf("Path{Path: %q, Level0: %q, Level1: %q, Level2: %q, Valid: %v, WildCard: %v}",
		p.path, p.lvl0, p.lvl1, p.lvl2, p.valid, p.wildcard,
	)
}

func init() {
	{
		f := func() reflect.Value {
			var o Path
			return reflect.ValueOf(&o)
		}
		rtypes.Factory.Add("AliCDBPath", f)
	}
}

func init() {
	// Streamer for AliCDBPath.
	rdict.Streamers.Add(rdict.NewCxxStreamerInfo("AliCDBPath", 1, 0xba0a3f48, []rbytes.StreamerElement{
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
		&rdict.StreamerString{StreamerElement: rdict.Element{
			Name:   *rbase.NewNamed("fPath", "detector pathname (Detector/DBType/SpecType)"),
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
		&rdict.StreamerString{StreamerElement: rdict.Element{
			Name:   *rbase.NewNamed("fLevel0", "level0 name (ex. detector: ZDC, TPC...)"),
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
		&rdict.StreamerString{StreamerElement: rdict.Element{
			Name:   *rbase.NewNamed("fLevel1", "level1 name (ex. DB type, Calib, Align)"),
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
		&rdict.StreamerString{StreamerElement: rdict.Element{
			Name:   *rbase.NewNamed("fLevel2", "level2 name (ex. DetSpecType, pedestals, gain...)"),
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
		&rdict.StreamerBasicType{StreamerElement: rdict.Element{
			Name:   *rbase.NewNamed("fIsValid", "validity flag"),
			Type:   rmeta.Bool,
			Size:   1,
			ArrLen: 0,
			ArrDim: 0,
			MaxIdx: [5]int32{0, 0, 0, 0, 0},
			Offset: 0,
			EName:  "Bool_t",
			XMin:   0.000000,
			XMax:   0.000000,
			Factor: 0.000000,
		}.New()},
		&rdict.StreamerBasicType{StreamerElement: rdict.Element{
			Name:   *rbase.NewNamed("fIsWildcard", "wildcard flag"),
			Type:   rmeta.Bool,
			Size:   1,
			ArrLen: 0,
			ArrDim: 0,
			MaxIdx: [5]int32{0, 0, 0, 0, 0},
			Offset: 0,
			EName:  "Bool_t",
			XMin:   0.000000,
			XMax:   0.000000,
			Factor: 0.000000,
		}.New()},
	}))
}

var (
	_ root.Object        = (*Path)(nil)
	_ rbytes.Marshaler   = (*Path)(nil)
	_ rbytes.Unmarshaler = (*Path)(nil)
)
