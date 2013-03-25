// Copyright 2012 Ruben Pollan <meskio@sindominio.net>
// Use of this source code is governed by a LGPL licence
// version 3 or later that can be found in the LICENSE file.

package epubgo

import "testing"

func TestIterator(t *testing.T) {
	f, _ := Open(book_path)
	defer f.Close()

	it := f.Navigation()
	if it.HasChildren() {
		t.Errorf("it.HasChildren() not behaving as expected")
		return
	}
}
