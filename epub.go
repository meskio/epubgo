// Copyright 2012 Ruben Pollan <meskio@sindominio.net>
// Use of this source code is governed by a LGPL licence
// version 3 or later that can be found in the LICENSE file.

package epub

import (
	"archive/zip"
)

type Epub struct {
	file     *zip.ReadCloser
	Metadata MData
}

func Open(path string) (e *Epub, err error) {
	e = new(Epub)
	e.file, err = zip.OpenReader(path)
	if err != nil {
		return
	}

	e.Metadata, err = parseMetadata(e.file)
	return
}

func (e Epub) Close() {
	e.file.Close()
}
