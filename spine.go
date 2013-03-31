// Copyright 2012 Ruben Pollan <meskio@sindominio.net>
// Use of this source code is governed by a LGPL licence
// version 3 or later that can be found in the LICENSE file.

package epubgo

import (
	"errors"
	"io"
)

// Iterator on the epub pages spine
//
// With it is possible to navigate throw the pages of the epub.
type SpineIterator struct {
	opf   *xmlOPF
	index int
	epub  *Epub
}

func newSpineIterator(epub *Epub) (*SpineIterator, error) {
	if epub.opf.spineLength() == 0 {
		return nil, errors.New("Spine is empty")
	}
	var spine SpineIterator
	spine.epub = epub
	spine.opf = epub.opf
	spine.index = 0
	return &spine, nil
}

// Is it the first element of the book?
func (spine SpineIterator) IsFirst() bool {
	return spine.index == 0
}

// Is it the last element of the book?
func (spine SpineIterator) IsLast() bool {
	return spine.index == spine.opf.spineLength()-1
}

// Advance the iterator to the next element on the spine
//
// Returns an error if is the last
func (spine *SpineIterator) Next() error {
	if spine.IsLast() {
		return errors.New("It is the last entry")
	}
	spine.index++
	return nil
}

// Step back the iterator to the previous element on the spine
//
// Returns an error if is the first
func (spine *SpineIterator) Previous() error {
	if spine.IsFirst() {
		return errors.New("It is the first entry")
	}
	spine.index--
	return nil
}

// Open the file of the iterator
func (spine SpineIterator) Open() (io.ReadCloser, error) {
	url := spine.Url()
	return spine.epub.OpenFile(url)
}

// Get the url of the item on the iterator
func (spine SpineIterator) Url() string {
	return spine.opf.spineUrl(spine.index)
}
