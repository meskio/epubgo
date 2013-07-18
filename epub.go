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

// Open opens an existing epub
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

// Load loads an epub from an io.ReaderAt
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
	ncxPath := e.opf.ncxPath()
	if ncxPath == "" {
		return errors.New("There is no NCX file on the epub")
	}
	ncx, err := e.OpenFile(ncxPath)
	if err != nil {
		return errors.New("Can't open the NCX file")
	}
	defer ncx.Close()
	e.ncx, err = parseNCX(ncx)
	return
}

// Close closes the epub file
func (e Epub) Close() {
	if e.file != nil {
		e.file.Close()
	}
}

// OpenFile opens a file inside the epub
func (e Epub) OpenFile(name string) (io.ReadCloser, error) {
	return openFile(e.zip, e.rootPath+name)
}

// OpenFileId opens a file from it's id
//
// The id of the files often appears on metadata fields
func (e Epub) OpenFileId(id string) (io.ReadCloser, error) {
	path := e.opf.filePath(id)
	return openFile(e.zip, e.rootPath+path)
}

// Navigation returns a navigation iterator
func (e Epub) Navigation() (*NavigationIterator, error) {
	return newNavigationIterator(e.ncx.navMap())
}

// Spine returns a spine iterator
func (e Epub) Spine() (*SpineIterator, error) {
	return newSpineIterator(&e)
}

// Metadata returns the values of a metadata field
//
// The valid field names are:
//    title, language, identifier, creator, subject, description, publisher,
//    contributor, date, type, format, source, relation, coverage, rights, meta
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

// MetadataFields retunrs the list of metadata fields pressent on the current epub
func (e Epub) MetadataFields() []string {
	fields := make([]string, len(e.metadata))
	i := 0
	for k, _ := range e.metadata {
		fields[i] = k
		i++
	}
	return fields
}

// MetadataAttr returns the metadata attributes
//
// Returns: an array of maps of each attribute and it's value.
// The array has the fields on the same order than the Metadata method.
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
