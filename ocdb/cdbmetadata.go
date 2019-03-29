// Copyright 2019 The Alice-Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ocdb

import (
	"fmt"
	"io"
	"reflect"

	"go-hep.org/x/hep/groot/rbase"
	"go-hep.org/x/hep/groot/rbytes"
	"go-hep.org/x/hep/groot/rcont"
	"go-hep.org/x/hep/groot/rdict"
	"go-hep.org/x/hep/groot/rmeta"
	"go-hep.org/x/hep/groot/root"
	"go-hep.org/x/hep/groot/rtypes"
)

// MetaData stores optional metadata associated with an entry in an OCDB file.
type MetaData struct {
	base    rbase.Object `groot:"BASE-TObject"`     // base class
	class   string       `groot:"fObjectClassName"` // object's class name
	resp    string       `groot:"fResponsible"`     // object's responsible person
	beam    uint32       `groot:"fBeamPeriod"`      // beam period
	vers    string       `groot:"fAliRootVersion"`  // AliRoot version
	comment string       `groot:"fComment"`         // extra comments
	props   rcont.Map    `groot:"fProperties"`      // list of object specific properties
}

func (*MetaData) Class() string   { return "AliCDBMetaData" }
func (*MetaData) RVersion() int16 { return 1 }

// MarshalROOT implements rbytes.Marshaler
func (meta *MetaData) MarshalROOT(w *rbytes.WBuffer) (int, error) {
	if w.Err() != nil {
		return 0, w.Err()
	}

	pos := w.WriteVersion(meta.RVersion())

	meta.base.MarshalROOT(w)
	w.WriteString(meta.class)
	w.WriteString(meta.resp)
	w.WriteU32(meta.beam)
	w.WriteString(meta.vers)
	w.WriteString(meta.comment)
	meta.props.MarshalROOT(w)

	return w.SetByteCount(pos, meta.Class())
}

// ROOTUnmarshaler is the interface implemented by an object that can
// unmarshal itself from a ROOT buffer
func (meta *MetaData) UnmarshalROOT(r *rbytes.RBuffer) error {
	if r.Err() != nil {
		return r.Err()
	}

	start := r.Pos()
	_, pos, bcnt := r.ReadVersion(meta.Class())

	if err := meta.base.UnmarshalROOT(r); err != nil {
		return err
	}

	meta.class = r.ReadString()
	meta.resp = r.ReadString()
	meta.beam = r.ReadU32()
	meta.vers = r.ReadString()
	meta.comment = r.ReadString()

	if err := meta.props.UnmarshalROOT(r); err != nil {
		return err
	}

	r.CheckByteCount(pos, bcnt, start, meta.Class())
	return r.Err()
}

func (meta *MetaData) Display(w io.Writer) {
	fmt.Fprintf(w, "Class: %q\nResponsible: %q\nBeamPeriod: %d\nAliRoot Version: %q\nComment: %q\nProperties: %d\n",
		meta.class, meta.resp, meta.beam, meta.vers, meta.comment, len(meta.props.Table()),
	)
	for k, v := range meta.props.Table() {
		fmt.Fprintf(w, "  key: %v\n  val: %v\n", k, v)
	}
}

func init() {
	{
		f := func() reflect.Value {
			var o MetaData
			return reflect.ValueOf(&o)
		}
		rtypes.Factory.Add("AliCDBMetaData", f)
	}
}

func init() {
	// Streamer for AliCDBMetaData.
	rdict.Streamers.Add(rdict.NewCxxStreamerInfo("AliCDBMetaData", 1, 0x746cbdf0, []rbytes.StreamerElement{
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
			Name:   *rbase.NewNamed("fObjectClassName", "object's class name"),
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
			Name:   *rbase.NewNamed("fResponsible", "object's responsible person"),
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
			Name:   *rbase.NewNamed("fBeamPeriod", "beam period"),
			Type:   rmeta.UInt,
			Size:   4,
			ArrLen: 0,
			ArrDim: 0,
			MaxIdx: [5]int32{0, 0, 0, 0, 0},
			Offset: 0,
			EName:  "UInt_t",
			XMin:   0.000000,
			XMax:   0.000000,
			Factor: 0.000000,
		}.New()},
		&rdict.StreamerString{StreamerElement: rdict.Element{
			Name:   *rbase.NewNamed("fAliRootVersion", "AliRoot version"),
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
			Name:   *rbase.NewNamed("fComment", "extra comments"),
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
		&rdict.StreamerObject{StreamerElement: rdict.Element{
			Name:   *rbase.NewNamed("fProperties", "list of object specific properties"),
			Type:   rmeta.Object,
			Size:   56,
			ArrLen: 0,
			ArrDim: 0,
			MaxIdx: [5]int32{0, 0, 0, 0, 0},
			Offset: 0,
			EName:  "TMap",
			XMin:   0.000000,
			XMax:   0.000000,
			Factor: 0.000000,
		}.New()},
	}))
}

var (
	_ root.Object        = (*MetaData)(nil)
	_ rbytes.Marshaler   = (*MetaData)(nil)
	_ rbytes.Unmarshaler = (*MetaData)(nil)
)
