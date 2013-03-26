// Copyright 2012 Ruben Pollan <meskio@sindominio.net>
// Use of this source code is governed by a LGPL licence
// version 3 or later that can be found in the LICENSE file.

package epubgo

import (
	"encoding/xml"
	"io"
)

type xmlNCX struct {
	NavMap []navpoint `xml:"navMap>navPoint"`
}
type navpoint struct {
	Text     string     `xml:"navLabel>text"`
	Content  string     `xml:"content>src,attr"`
	NavPoint []navpoint `xml:"navPoint"`
}

func parseNCX(ncx io.Reader) (*xmlNCX, error) {
	decoder := xml.NewDecoder(ncx)
	var n xmlNCX
	err := decoder.Decode(&n)
	if err != nil {
		return nil, err
	}

	return &n, nil
}

func (ncx xmlNCX) navMap() []navpoint {
	return ncx.NavMap
}

func (point navpoint) Title() string {
	return point.Text
}

func (point navpoint) Children() []navpoint {
	return point.NavPoint
}
