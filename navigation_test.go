// Copyright 2012 Ruben Pollan <meskio@sindominio.net>
// Use of this source code is governed by a LGPL licence
// version 3 or later that can be found in the LICENSE file.

package epubgo

import "testing"

const (
	first_title = "A DOG'S TALE, By Mark Twain"
	first_url   = "@public@vhost@g@gutenberg@html@files@3174@3174-h@3174-h-0.htm.html#pgepubid00000"
	child_title = "Frontpiece"
)

func TestIterator(t *testing.T) {
	f, _ := Open(book_path)
	defer f.Close()

	it := f.Navigation()
	if it.HasChildren() {
		t.Errorf("it.HasChildren() not behaving as expected")
	}
	if it.HasParents() {
		t.Errorf("it.HasParents() not behaving as expected")
	}
	if !it.IsFirst() {
		t.Errorf("it.IsFirst() not behaving as expected")
	}
	if it.IsLast() {
		t.Errorf("it.IsLast() not behaving as expected")
	}
}

func TestTitle(t *testing.T) {
	f, _ := Open(book_path)
	defer f.Close()

	it := f.Navigation()
	if it.Title() != first_title {
		t.Errorf("it.Title() return: %v when was expected: %v", it.Title(), first_title)
	}
}

func TestUrl(t *testing.T) {
	f, _ := Open(book_path)
	defer f.Close()

	it := f.Navigation()
	if it.Url() != first_url {
		t.Errorf("it.Url() return: %v when was expected: %v", it.Url(), first_url)
	}
}

func TestDepth(t *testing.T) {
	f, _ := Open(book_path)
	defer f.Close()

	it := f.Navigation()
	if it.In() == nil {
		t.Errorf("it.In() din't return an error whithout having children")
	}
	if err := it.Next(); err != nil {
		t.Errorf("it.Next() return an error: %v", err)
	}
	if err := it.Next(); err != nil {
		t.Errorf("it.Next() return an error: %v", err)
	}
	if !it.IsLast() {
		t.Errorf("it.IsLast() not behaving as expected")
	}
	if err := it.In(); err != nil {
		t.Errorf("it.In() return an error: %v", err)
	}
	if it.Previous() == nil {
		t.Errorf("it.Previous() din't return an error being the first")
	}
	if err := it.Next(); err != nil {
		t.Errorf("it.Next() return an error: %v", err)
	}
	if err := it.Previous(); err != nil {
		t.Errorf("it.Previous() return an error: %v", err)
	}
	if it.Title() != child_title {
		t.Errorf("it.Title() return: %v when was expected: %v", it.Title(), child_title)
	}
	if err := it.Out(); err != nil {
		t.Errorf("it.Out() return an error: %v", err)
	}
	if err := it.Previous(); err != nil {
		t.Errorf("it.Next() return an error: %v", err)
	}
	if err := it.Previous(); err != nil {
		t.Errorf("it.Next() return an error: %v", err)
	}
	if it.Title() != first_title {
		t.Errorf("it.Title() return: %v when was expected: %v", it.Title(), first_title)
	}
}
