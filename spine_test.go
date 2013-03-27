// Copyright 2012 Ruben Pollan <meskio@sindominio.net>
// Use of this source code is governed by a LGPL licence
// version 3 or later that can be found in the LICENSE file.

package epubgo

import "testing"

func TestFirst(t *testing.T) {
	f, _ := Open(book_path)
	defer f.Close()

	it := f.Spine()
	if !it.IsFirst() {
		t.Errorf("it.IsFirst() not behaving as expected")
	}
	if it.IsLast() {
		t.Errorf("it.IsLast() not behaving as expected")
	}
}

func TestLast(t *testing.T) {
	f, _ := Open(book_path)
	defer f.Close()

	it := f.Spine()
	if err := it.Next(); err != nil {
		t.Errorf("it.Next() return an error: %v", err)
	}
	if it.IsFirst() {
		t.Errorf("it.IsFirst() not behaving as expected")
	}
	if !it.IsLast() {
		t.Errorf("it.IsLast() not behaving as expected")
	}
}
