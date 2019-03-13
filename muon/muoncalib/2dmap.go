// Copyright 2019 The Alice-Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package muoncalib

import (
	"fmt"
	"sort"

	"go-hep.org/x/hep/groot/root"
)

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
