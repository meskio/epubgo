// Copyright 2012 Ruben Pollan <meskio@sindominio.net>
// Use of this source code is governed by a LGPL licence
// version 3 or later that can be found in the LICENSE file.

package epub

import (
	"archive/zip"
)

// Epub holds all the data of the ebook
type Epub struct {
	file     *zip.ReadCloser
	Metadata MData
}

// Open an existing epub
func Open(path string) (e *Epub, err error) {
	e = new(Epub)
	e.file, err = zip.OpenReader(path)
	if err != nil {
		return
	}

	e.Metadata, err = parseMetadata(e.file)
	return
}

// Close the epub file
func (e Epub) Close() {
	e.file.Close()
}
