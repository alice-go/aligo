// Copyright 2019 The Alice-Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package muoncalib

import (
	"go-hep.org/x/hep/groot/rbase"
	"go-hep.org/x/hep/groot/rbytes"
	"go-hep.org/x/hep/groot/rcont"
)

type AliMUON2DMap struct {
	base  AliMUONVStore
	exmap *AliMpExMap `groot:"fMap"`
	opt   bool        `groot:"fOptimizeForDEManu"`
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

type AliMUONVStore struct {
	base rbase.Object
}

func (*AliMUONVStore) Class() string   { return "AliMUONVStore" }
func (*AliMUONVStore) RVersion() int16 { return 1 }

// MarshalROOT implements rbytes.Marshaler
func (o *AliMUONVStore) MarshalROOT(w *rbytes.WBuffer) (int, error) {
	if w.Err() != nil {
		return 0, w.Err()
	}

	pos := w.WriteVersion(o.RVersion())

	o.base.MarshalROOT(w)

	return w.SetByteCount(pos, o.Class())
}

// ROOTUnmarshaler is the interface implemented by an object that can
// unmarshal itself from a ROOT buffer
func (o *AliMUONVStore) UnmarshalROOT(r *rbytes.RBuffer) error {
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

type AliMpExMap struct {
	base rbase.Object
	objs rcont.ObjArray `groot:"fObjects"`
	keys rcont.ArrayL64 `groot:"fKeys"`
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

type AliMUONCalibParamND struct {
	base AliMUONVCalibParam
	dim  int32     `groot:"fDimension"`
	size int32     `groot:"fSize"`
	n    int32     `groot:"fN"`
	vs   []float64 `groot:"fValues"`
}

func (*AliMUONCalibParamND) Class() string   { return "AliMUONCalibParamND" }
func (*AliMUONCalibParamND) RVersion() int16 { return 1 }

// MarshalROOT implements rbytes.Marshaler
func (o *AliMUONCalibParamND) MarshalROOT(w *rbytes.WBuffer) (int, error) {
	if w.Err() != nil {
		return 0, w.Err()
	}

	pos := w.WriteVersion(o.RVersion())

	o.base.MarshalROOT(w)
	w.WriteI32(o.dim)
	w.WriteI32(o.size)
	w.WriteI32(o.n)
	w.WriteI8(1) // FIXME(sbinet)
	w.WriteFastArrayF64(o.vs)

	return w.SetByteCount(pos, o.Class())
}

// ROOTUnmarshaler is the interface implemented by an object that can
// unmarshal itself from a ROOT buffer
func (o *AliMUONCalibParamND) UnmarshalROOT(r *rbytes.RBuffer) error {
	if r.Err() != nil {
		return r.Err()
	}

	start := r.Pos()
	_, pos, bcnt := r.ReadVersion(o.Class())

	if err := o.base.UnmarshalROOT(r); err != nil {
		return err
	}

	o.dim = r.ReadI32()
	o.size = r.ReadI32()
	o.n = r.ReadI32()
	_ = r.ReadI8() // FIXME(sbinet)
	o.vs = r.ReadFastArrayF64(int(o.n))

	r.CheckByteCount(pos, bcnt, start, o.Class())
	return r.Err()
}

type AliMUONVCalibParam struct {
	base rbase.Object
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
