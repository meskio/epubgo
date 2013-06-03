// Copyright 2012 Ruben Pollan <meskio@sindominio.net>
// Use of this source code is governed by a LGPL licence
// version 3 or later that can be found in the LICENSE file.

package epubgo

import "testing"

import (
	"os"
)

const (
	encodingOpf = "testdata/encoding_err.opf"
)

func TestEncodingError(t *testing.T) {
	file, _ := os.Open(encodingOpf)
	defer file.Close()

	_, err := parseOPF(file)
	if err != nil {
		t.Errorf("parseOpf(%v) with encoding problems return an error: %v", encodingOpf, err)
	}
}
