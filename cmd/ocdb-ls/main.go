// Copyright 2019 The Alice-Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"os"

	_ "github.com/alice-go/aligo/muon/muoncalib"
	"github.com/alice-go/aligo/ocdb"
	"go-hep.org/x/hep/groot"
	_ "go-hep.org/x/hep/groot/ztypes"
)

func main() {
	flag.Parse()

	f, err := groot.Open(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for _, k := range f.Keys() {
		log.Printf("key: %v", k.Name())

		o, err := f.Get(k.Name())
		if err != nil {
			log.Fatal(err)
		}

		v := o.(*ocdb.Entry)
		v.Display(os.Stdout)
	}
}
