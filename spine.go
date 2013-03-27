// Copyright 2012 Ruben Pollan <meskio@sindominio.net>
// Use of this source code is governed by a LGPL licence
// version 3 or later that can be found in the LICENSE file.

package epubgo

import (
	"errors"
)

// Iterator on the epub pages spine
//
// With it is possible to navigate throw the pages of the epub.
type SpineIterator struct {
	opf   *xmlOPF
	index int
}

func newSpineIterator(opf *xmlOPF) *SpineIterator {
	var spine SpineIterator
	spine.opf = opf
	spine.index = 0
	return &spine
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
