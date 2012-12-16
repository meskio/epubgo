// Copyright 2012 Ruben Pollan <meskio@sindominio.net>
// Use of this source code is governed by a LGPL licence
// version 3 or later that can be found in the LICENSE file.

package epub

import (
	"archive/zip"
	"errors"
)

type element struct {
	content string
	attr    map[string]string
}
type mdata map[string][]element

// Epub holds all the data of the ebook
type Epub struct {
	file     *zip.ReadCloser
	metadata mdata
}

// Open an existing epub
func Open(path string) (e *Epub, err error) {
	e = new(Epub)
	e.file, err = zip.OpenReader(path)
	if err != nil {
		return
	}

	e.metadata, err = parseMetadata(e.file)
	return
}

// Close the epub file
func (e Epub) Close() {
	e.file.Close()
}

func (e Epub) Metadata(field string) ([]string, error) {
	elem, ok := e.metadata[field]
	if ok {
		cont := make([]string, len(elem))
		for i, e := range elem {
			cont[i] = e.content
		}
		return cont, nil
	}

	return nil, errors.New("Field " + field + " don't exists")
}

/*func (m MData) Len(field string) int {
	elem, ok := m[field]
	if ok {
		return len(elem)
	}

	return 0
}*/
