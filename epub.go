// Copyright 2012 Ruben Pollan <meskio@sindominio.net>
// Use of this source code is governed by a LGPL licence
// version 3 or later that can be found in the LICENSE file.

package epubgo

import (
	"archive/zip"
	"errors"
	"io"
	"os"
)

// Epub holds all the data of the ebook
type Epub struct {
	file     *os.File
	zip      *zip.Reader
	rootPath string
	metadata mdata
	opf      *xmlOPF
	ncx      *xmlNCX
}

type mdata map[string][]mdataElement
type mdataElement struct {
	content string
	attr    map[string]string
}

// Open an existing epub
func Open(path string) (e *Epub, err error) {
	e = new(Epub)
	e.file, err = os.Open(path)
	if err != nil {
		return
	}
	fileInfo, err := e.file.Stat()
	if err != nil {
		return
	}
	err = e.load(e.file, fileInfo.Size())
	return
}

// Load an eupb from an io.ReaderAt
func Load(r io.ReaderAt, size int64) (e *Epub, err error) {
	e = new(Epub)
	e.file = nil
	err = e.load(r, size)
	return
}

func (e *Epub) load(r io.ReaderAt, size int64) (err error) {
	e.zip, err = zip.NewReader(r, size)
	if err != nil {
		return
	}

	e.rootPath, err = getRootPath(e.zip)
	if err != nil {
		return
	}

	return e.parseFiles()
}

func (e *Epub) parseFiles() (err error) {
	opfFile, err := openOPF(e.zip)
	if err != nil {
		return
	}
	defer opfFile.Close()
	e.opf, err = parseOPF(opfFile)
	if err != nil {
		return
	}

	e.metadata = e.opf.toMData()
	if err != nil {
		return
	}
	ncx, err := e.OpenFile(e.opf.ncxPath())
	if err != nil {
		return
	}
	defer ncx.Close()
	e.ncx, err = parseNCX(ncx)
	return
}

// Close the epub file
func (e Epub) Close() {
	if e.file != nil {
		e.file.Close()
	}
}

// Open a file inside the epub
func (e Epub) OpenFile(name string) (io.ReadCloser, error) {
	return openFile(e.zip, e.rootPath+name)
}

// Get a navigation iterator
func (e Epub) Navigation() *NavigationIterator {
	return newNavigationIterator(e.ncx.navMap())
}

// Get the values of a metadata field
//
// The valid field names are:
//    Title, Language, Identifier, Creator, Subject, Description, Publisher, 
//    Contributor, Date, Type, Format, Source, Relation, Coverage, Rights
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

// Get the list of metadata fields
func (e Epub) MetadataFields() []string {
	fields := make([]string, len(e.metadata))
	i := 0
	for k, _ := range e.metadata {
		fields[i] = k
		i++
	}
	return fields
}

// Get the metadata attributes
//
// The array  has the fields on the smae order than the Metadata method
func (e Epub) MetadataAttr(field string) ([]map[string]string, error) {
	elem, ok := e.metadata[field]
	if ok {
		attr := make([]map[string]string, len(elem))
		for i, e := range elem {
			attr[i] = e.attr
		}
		return attr, nil
	}

	return nil, errors.New("Field " + field + " don't exists")
}
