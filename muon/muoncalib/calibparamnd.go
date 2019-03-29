// Copyright 2019 The Alice-Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package muoncalib

import (
	"fmt"
	"io"
	"math"
	"reflect"
	"strings"

	"go-hep.org/x/hep/groot/rbase"
	"go-hep.org/x/hep/groot/rbytes"
	"go-hep.org/x/hep/groot/rdict"
	"go-hep.org/x/hep/groot/rmeta"
	"go-hep.org/x/hep/groot/root"
	"go-hep.org/x/hep/groot/rtypes"
)

type AliMUONCalibParamND struct {
	base AliMUONVCalibParam `groot:"BASE-AliMUONVCalibParam"` // base class
	dim  int32              `groot:"fDimension"`              // /< dimension of this object
	size int32              `groot:"fSize"`                   // /< The number of double tuples we hold
	n    int32              `groot:"fN"`                      // /< The total number of floats we hold (fDimension*fSize)
	vs   []float64          `groot:"fValues,meta=[fN]"`       // The values array
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

func (c *AliMUONCalibParamND) index(i, j int) int {
	return i + int(c.size)*j
}

func (c *AliMUONCalibParamND) ID0() uint32 {
	return c.base.base.ID & 0xFFFF
}

func (c *AliMUONCalibParamND) ID1() uint32 {
	return (c.base.base.ID & 0xFFFF0000) >> 16
}

func (c *AliMUONCalibParamND) Value(i, j int) float64 {
	return c.vs[c.index(i, j)]
}

func (c *AliMUONCalibParamND) MeanAndSigma(dim int) (float64, float64) {
	mean := 0.0
	v2 := 0.0
	n := int(c.size)
	for i := 0; i < n; i++ {
		v := c.Value(i, dim)
		mean += v
		v2 += v * v
	}
	mean /= float64(n)
	sigma := 0.0
	if n > 1 {
		sigma = math.Sqrt((v2 - float64(n)*mean*mean) / (float64(n) - 1))
	}
	return mean, sigma
}

func (c *AliMUONCalibParamND) Print(w io.Writer, opt string) {
	fmt.Fprintf(w, "AliMUONCalibParamND Id=(%d,%d) Size=%d Dimension=%d\n",
		c.ID0(), c.ID1(), c.size, c.dim)
	opt = strings.ToUpper(opt)
	if strings.Contains(opt, "FULL") {
		for i := 0; i < int(c.size); i++ {
			fmt.Fprintf(w, "CH %3d", i)
			for j := 0; j < int(c.dim); j++ {
				fmt.Fprintf(w, " %g", c.Value(i, j))
			}
			fmt.Fprint(w, "\n")
		}
	}
	if strings.Contains(opt, "MEAN") {
		var j int
		fmt.Sscanf(opt, "MEAN%d", &j)
		mean, sigma := c.MeanAndSigma(j)
		fmt.Fprintf(w, " Mean(j=%d)=%g Sigma(j=%d)=%g\n", j, mean, j, sigma)
	}
}

func init() {
	{
		f := func() reflect.Value {
			var o AliMUONCalibParamND
			return reflect.ValueOf(&o)
		}
		rtypes.Factory.Add("AliMUONCalibParamND", f)
	}
}

func init() {
	// Streamer for AliMUONCalibParamND.
	rdict.Streamers.Add(rdict.NewCxxStreamerInfo("AliMUONCalibParamND", 1, 0xb1a7e64a, []rbytes.StreamerElement{
		rdict.NewStreamerBase(rdict.Element{
			Name:   *rbase.NewNamed("AliMUONVCalibParam", "Base class for a calibration data holder (usually for 64 channels)"),
			Type:   rmeta.Base,
			Size:   0,
			ArrLen: 0,
			ArrDim: 0,
			MaxIdx: [5]int32{0, -71169467, 0, 0, 0},
			Offset: 0,
			EName:  "BASE",
			XMin:   0.000000,
			XMax:   0.000000,
			Factor: 0.000000,
		}.New(), 0),
		&rdict.StreamerBasicType{StreamerElement: rdict.Element{
			Name:   *rbase.NewNamed("fDimension", "/< dimension of this object"),
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
			Name:   *rbase.NewNamed("fSize", "/< The number of double tuples we hold"),
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
			Name:   *rbase.NewNamed("fN", "/< The total number of floats we hold (fDimension*fSize)"),
			Type:   rmeta.Counter,
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
		rdict.NewStreamerBasicPointer(rdict.Element{
			Name:   *rbase.NewNamed("fValues", "[fN] The values array"),
			Type:   48,
			Size:   8,
			ArrLen: 0,
			ArrDim: 0,
			MaxIdx: [5]int32{0, 0, 0, 0, 0},
			Offset: 0,
			EName:  "Double_t*",
			XMin:   0.000000,
			XMax:   0.000000,
			Factor: 0.000000,
		}.New(), 1, "fN", "AliMUONCalibParamND"),
	}))
}

var (
	_ root.Object        = (*AliMUONCalibParamND)(nil)
	_ rbytes.Marshaler   = (*AliMUONCalibParamND)(nil)
	_ rbytes.Unmarshaler = (*AliMUONCalibParamND)(nil)
)
