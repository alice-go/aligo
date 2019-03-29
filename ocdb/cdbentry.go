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
	"go-hep.org/x/hep/groot/rdict"
	"go-hep.org/x/hep/groot/rmeta"
	"go-hep.org/x/hep/groot/root"
	"go-hep.org/x/hep/groot/rtypes"
)

// Entry represents a single entry in an OCDB data file.
type Entry struct {
	base  rbase.Object `groot:"BASE-TObject"` // base class
	obj   root.Object  `groot:"fObject"`      // object
	id    ID           `groot:"fId"`          // entry ID
	meta  *MetaData    `groot:"fMetaData"`    // metaData
	owner bool         `groot:"fIsOwner"`     // ownership flag
}

func (*Entry) Class() string   { return "AliCDBEntry" }
func (*Entry) RVersion() int16 { return 1 }

func (entry *Entry) Object() root.Object { return entry.obj }
func (entry *Entry) Id() ID              { return entry.id }

func (entry *Entry) Display(w io.Writer) {
	fmt.Fprintf(w, `=== Entry ===
ID: %v
Owner: %v
`,
		entry.id, entry.owner,
	)
	if entry.meta != nil {
		fmt.Fprintf(w, "MetaData:\n")
		entry.meta.Display(w)
	}
	if entry.obj != nil {
		fmt.Fprintf(w, "Object: %T\n%v\n===\n", entry.obj, entry.obj)
	}
}

// MarshalROOT implements rbytes.Marshaler
func (entry *Entry) MarshalROOT(w *rbytes.WBuffer) (int, error) {
	if w.Err() != nil {
		return 0, w.Err()
	}

	pos := w.WriteVersion(entry.RVersion())

	entry.base.MarshalROOT(w)
	w.WriteObjectAny(entry.obj)
	entry.id.MarshalROOT(w)
	w.WriteObjectAny(entry.meta)
	w.WriteBool(entry.owner)

	return w.SetByteCount(pos, entry.Class())
}

// ROOTUnmarshaler is the interface implemented by an object that can
// unmarshal itself from a ROOT buffer
func (o *Entry) UnmarshalROOT(r *rbytes.RBuffer) error {
	if r.Err() != nil {
		return r.Err()
	}

	start := r.Pos()
	_, pos, bcnt := r.ReadVersion(o.Class())

	if err := o.base.UnmarshalROOT(r); err != nil {
		return err
	}

	o.obj = r.ReadObjectAny()
	if err := o.id.UnmarshalROOT(r); err != nil {
		return err
	}
	o.meta = nil
	if obj := r.ReadObjectAny(); obj != nil {
		o.meta = obj.(*MetaData)
	}
	o.owner = r.ReadBool()

	r.CheckByteCount(pos, bcnt, start, o.Class())
	return r.Err()
}

func init() {
	{
		f := func() reflect.Value {
			var o Entry
			return reflect.ValueOf(&o)
		}
		rtypes.Factory.Add("AliCDBEntry", f)
	}
}

func init() {
	// Streamer for AliCDBEntry.
	rdict.Streamers.Add(rdict.NewCxxStreamerInfo("AliCDBEntry", 1, 0x9d5eed3c, []rbytes.StreamerElement{
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
		&rdict.StreamerObjectPointer{StreamerElement: rdict.Element{
			Name:   *rbase.NewNamed("fObject", "object"),
			Type:   rmeta.ObjectP,
			Size:   8,
			ArrLen: 0,
			ArrDim: 0,
			MaxIdx: [5]int32{0, 0, 0, 0, 0},
			Offset: 0,
			EName:  "TObject*",
			XMin:   0.000000,
			XMax:   0.000000,
			Factor: 0.000000,
		}.New()},
		&rdict.StreamerObject{StreamerElement: rdict.Element{
			Name:   *rbase.NewNamed("fId", "entry ID"),
			Type:   rmeta.Object,
			Size:   192,
			ArrLen: 0,
			ArrDim: 0,
			MaxIdx: [5]int32{0, 0, 0, 0, 0},
			Offset: 0,
			EName:  "AliCDBId",
			XMin:   0.000000,
			XMax:   0.000000,
			Factor: 0.000000,
		}.New()},
		&rdict.StreamerObjectPointer{StreamerElement: rdict.Element{
			Name:   *rbase.NewNamed("fMetaData", "metaData"),
			Type:   rmeta.ObjectP,
			Size:   8,
			ArrLen: 0,
			ArrDim: 0,
			MaxIdx: [5]int32{0, 0, 0, 0, 0},
			Offset: 0,
			EName:  "AliCDBMetaData*",
			XMin:   0.000000,
			XMax:   0.000000,
			Factor: 0.000000,
		}.New()},
		&rdict.StreamerBasicType{StreamerElement: rdict.Element{
			Name:   *rbase.NewNamed("fIsOwner", "ownership flag"),
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
	_ root.Object        = (*Entry)(nil)
	_ rbytes.Marshaler   = (*Entry)(nil)
	_ rbytes.Unmarshaler = (*Entry)(nil)
)
