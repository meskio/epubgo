// Copyright 2012 Ruben Pollan <meskio@sindominio.net>
// Use of this source code is governed by a LGPL licence
// version 3 or later that can be found in the LICENSE file.

package epubgo

import "testing"

import (
	"os"
)

const (
	nbspNCX = "testdata/nbsp.ncx"
)

func TestNbspNcx(t *testing.T) {
	file, _ := os.Open(nbspNCX)
	defer file.Close()

	_, err := parseNCX(file)
	if err != nil {
		t.Errorf("parseNCX(%v) with encoding problems return an error: %v", nbspNCX, err)
	}
}
