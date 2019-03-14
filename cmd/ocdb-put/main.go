// Copyright 2019 The Alice-Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	_ "github.com/alice-go/aligo/muon/muoncalib"
	"github.com/alice-go/aligo/ocdb"
	"go-hep.org/x/hep/groot"
)

var (
	ccdb   string
	srcdir string
	dest   string
	dry    bool
	limit  int
)

func dumpRequest(r *http.Request) {
	output, err := httputil.DumpRequest(r, false)
	if err != nil {
		fmt.Println("Error dumping request:", err)
		return
	}
	fmt.Println(string(output))
}

func dumpResponse(r *http.Response) {
	output, err := httputil.DumpResponse(r, false)
	if err != nil {
		fmt.Println("Error dumping response:", err)
		return
	}
	fmt.Println(string(output))
}

func process(client *http.Client, path string, dest string, dry bool) {
	f, err := groot.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	key := "AliCDBEntry"
	o, err := f.Get(key)
	if err != nil {
		log.Fatalf("Could not get key %s from file %s\n", key, path)
	}
	v := o.(*ocdb.Entry)
	// using run range as timestamp range for the moment
	// FIXME: should read the corresponding GRP/GRP/Data object to
	// get the run->timestamp relationship and use timestamps
	// as validity range for the put
	r1 := v.Id().Runs().First
	// r2 := v.Id().Runs().Last
	url := ccdb + "/" + dest + "/" + strconv.Itoa(int(r1))

	if dry {
		fmt.Printf("Would upload %s to %s\n", path, url)
		return
	}

	r, err := os.Open(path)
	if err != nil {
		log.Fatalf("Cannot open file %s", path)
	}

	var requestBody bytes.Buffer
	mpw := multipart.NewWriter(&requestBody)

	w, err := mpw.CreateFormFile("data", path)
	if err != nil {
		log.Fatalf("Cannot create form file %s", err.Error())
	}

	_, err = io.Copy(w, r)
	if err != nil {
		log.Fatalf("Cannot copy file to request body %s", err.Error())
	}
	mpw.Close()

	req, err := http.NewRequest(http.MethodPost, url, &requestBody)
	if err != nil {
		log.Fatalf("Could not create request %s", err.Error())
	}

	req.Header.Set("Content-Type", mpw.FormDataContentType())
	dumpRequest(req)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Request did not go well %s", err.Error())
	}
	defer resp.Body.Close()
	dumpResponse(resp)
}

func init() {
	flag.StringVar(&srcdir, "srcdir", "/Users/laurent/cernbox/ocdbs/2018/OCDB/MUON/Calib/OccupancyMap", "local source directory containing OCDB objects")
	flag.StringVar(&dest, "dest", "OccupancyMap/MUON", "where to upload objects found in srcdir to")
	flag.StringVar(&ccdb, "ccdb", "http://localhost:6464", "URL of CCDB endpoint")
	flag.IntVar(&limit, "limit", 0, "limit the number of files that will be transfered (0 means no limit)")
	flag.BoolVar(&dry, "dry", false, "only list what would happen without doing it")
}

func main() {
	flag.Parse()
	processed := 0
	client := &http.Client{Timeout: 2 * time.Second}

	err := filepath.Walk(srcdir, func(path string, info os.FileInfo, err error) error {
		if strings.HasPrefix(filepath.Base(path), "Run") &&
			filepath.Ext(path) == ".root" {
			if limit != 0 && processed == limit {
				return io.EOF
			}
			process(client, path, dest, dry)
			processed++
		}
		return nil
	})
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
}
