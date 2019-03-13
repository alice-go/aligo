// Copyright 2019 The Alice-Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package muoncalib

import (
	"bytes"
	"fmt"

	"go-hep.org/x/hep/groot/rcont"
)

func (exmap AliMpExMap) String() string {
	o := new(bytes.Buffer)
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
