// Copyright 2019 The Alice-Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package muoncalib

import (
	"fmt"
	"reflect"
	"sort"

	"go-hep.org/x/hep/groot/rbase"
	"go-hep.org/x/hep/groot/rbytes"
	"go-hep.org/x/hep/groot/rdict"
	"go-hep.org/x/hep/groot/rmeta"
	"go-hep.org/x/hep/groot/root"
	"go-hep.org/x/hep/groot/rtypes"
)

type AliMUON2DMap struct {
	base  AliMUONVStore `groot:"BASE-AliMUONVStore"` // base class
	exmap *AliMpExMap   `groot:"fMap"`               // /< Our internal map (an AliMpExMap of AliMpExMaps)
	opt   bool          `groot:"fOptimizeForDEManu"` // /< whether (i,j) pair is supposed to be (DetElemId,ManuId) (allow us to allocate right amount of memory, that's all it does.

}

func (*AliMUON2DMap) Class() string   { return "AliMUON2DMap" }
func (*AliMUON2DMap) RVersion() int16 { return 1 }

// MarshalROOT implements rbytes.Marshaler
func (o *AliMUON2DMap) MarshalROOT(w *rbytes.WBuffer) (int, error) {
	if w.Err() != nil {
		return 0, w.Err()
	}

	pos := w.WriteVersion(o.RVersion())

	o.base.MarshalROOT(w)
	w.WriteObjectAny(o.exmap)
	w.WriteBool(o.opt)

	return w.SetByteCount(pos, o.Class())
}

// ROOTUnmarshaler is the interface implemented by an object that can
// unmarshal itself from a ROOT buffer
func (o *AliMUON2DMap) UnmarshalROOT(r *rbytes.RBuffer) error {
	if r.Err() != nil {
		return r.Err()
	}

	start := r.Pos()
	_, pos, bcnt := r.ReadVersion(o.Class())

	if err := o.base.UnmarshalROOT(r); err != nil {
		return err
	}

	o.exmap = nil
	if obj := r.ReadObjectAny(); obj != nil {
		o.exmap = obj.(*AliMpExMap)
	}
	o.opt = r.ReadBool()

	r.CheckByteCount(pos, bcnt, start, o.Class())
	return r.Err()
}

type Manu struct {
	DeID int
	ID   int
}

func sortedManus(manus []Manu) []Manu {
	sort.Slice(manus, func(i, j int) bool {
		if manus[i].DeID == manus[j].DeID {
			return manus[i].ID < manus[j].ID
		}
		return manus[i].DeID < manus[j].DeID
	})
	return manus
}

func (m *AliMUON2DMap) ExMap() *AliMpExMap { return m.exmap }

func (m *AliMUON2DMap) GetObject(deid, manuid int) root.Object {
	objects := m.exmap.Objects()
	keys := m.exmap.Keys()
	for i := 0; i < objects.Len(); i++ {
		de := keys.At(i)
		if int(de) != deid {
			continue
		}
		om := objects.At(i).(*AliMpExMap)
		k := om.Keys()
		o := om.Objects()
		for j := 0; j < o.Len(); j++ {
			if int(k.At(j)) != manuid {
				continue
			}
			return o.At(j)
		}
	}
	return nil
}

func (m *AliMUON2DMap) GetManusForDE(deid int) []Manu {
	var manus []Manu
	objects := m.exmap.Objects()
	keys := m.exmap.Keys()
	for i := 0; i < objects.Len(); i++ {
		de := keys.At(i)
		if int(de) != deid {
			continue
		}
		om := objects.At(i).(*AliMpExMap)
		k := om.Keys()
		o := om.Objects()
		for j := 0; j < o.Len(); j++ {
			manus = append(manus, Manu{int(de), int(k.At(j))})
		}
	}
	return sortedManus(manus)
}

func (m *AliMUON2DMap) GetManus() []Manu {
	var manus []Manu
	objects := m.exmap.Objects()
	keys := m.exmap.Keys()
	for i := 0; i < objects.Len(); i++ {
		de := keys.At(i)
		manus = append(manus, m.GetManusForDE(int(de))...)
	}
	return sortedManus(manus)
}

func (m *AliMUON2DMap) String() string {
	return fmt.Sprintf("MUON2DMap{Opt: %v, Map: %v}", m.opt, *m.exmap)
}

func init() {
	{
		f := func() reflect.Value {
			var o AliMUON2DMap
			return reflect.ValueOf(&o)
		}
		rtypes.Factory.Add("AliMUON2DMap", f)
	}
}

func init() {
	// Streamer for AliMUON2DMap.
	rdict.Streamers.Add(rdict.NewCxxStreamerInfo("AliMUON2DMap", 2, 0x6cd64b39, []rbytes.StreamerElement{
		rdict.NewStreamerBase(rdict.Element{
			Name:   *rbase.NewNamed("AliMUONVStore", "Base class for a MUON data store"),
			Type:   rmeta.Base,
			Size:   0,
			ArrLen: 0,
			ArrDim: 0,
			MaxIdx: [5]int32{0, -1733838666, 0, 0, 0},
			Offset: 0,
			EName:  "BASE",
			XMin:   0.000000,
			XMax:   0.000000,
			Factor: 0.000000,
		}.New(), 1),
		&rdict.StreamerObjectPointer{StreamerElement: rdict.Element{
			Name:   *rbase.NewNamed("fMap", "/< Our internal map (an AliMpExMap of AliMpExMaps)"),
			Type:   rmeta.ObjectP,
			Size:   8,
			ArrLen: 0,
			ArrDim: 0,
			MaxIdx: [5]int32{0, 0, 0, 0, 0},
			Offset: 0,
			EName:  "AliMpExMap*",
			XMin:   0.000000,
			XMax:   0.000000,
			Factor: 0.000000,
		}.New()},
		&rdict.StreamerBasicType{StreamerElement: rdict.Element{
			Name:   *rbase.NewNamed("fOptimizeForDEManu", "/< whether (i,j) pair is supposed to be (DetElemId,ManuId) (allow us to allocate right amount of memory, that's all it does."),
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
	_ root.Object        = (*AliMUON2DMap)(nil)
	_ rbytes.Marshaler   = (*AliMUON2DMap)(nil)
	_ rbytes.Unmarshaler = (*AliMUON2DMap)(nil)
)
