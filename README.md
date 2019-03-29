# aligo

[![Build Status](https://travis-ci.org/alice-go/aligo.svg?branch=master)](https://travis-ci.org/alice-go/aligo)
[![GoDoc](https://godoc.org/github.com/alice-go/aligo?status.svg)](https://godoc.org/github.com/alice-go/aligo)
[![codecov](https://codecov.io/gh/alice-go/aligo/branch/master/graph/badge.svg)](https://codecov.io/gh/alice-go/aligo)

`aligo` is a WIP package to manipulate OCDB files from the [ALICE](https://aliceinfo.cern.ch) experiment, in [Go](https://golang.org).

Currently available are two small programs :

- [`ocdb-ls`](cmd/ocdb-ls) to dump the content of an OCDB (Root) file
- [`ocdb-put`](cmd/ocdb-put) to upload OCDB files to a CCDB instance
