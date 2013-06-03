// Copyright 2012 Ruben Pollan <meskio@sindominio.net>
// Use of this source code is governed by a LGPL licence
// version 3 or later that can be found in the LICENSE file.

package epubgo

import "testing"

const (
	firstTitle = "A DOG'S TALE, By Mark Twain"
	firstURL   = "@public@vhost@g@gutenberg@html@files@3174@3174-h@3174-h-0.htm.html#pgepubid00000"
	childTitle = "Frontpiece"
)

func TestIterator(t *testing.T) {
	f, _ := Open(bookPath)
	defer f.Close()

	it, err := f.Navigation()
	if err != nil {
		t.Errorf("epub.Navigation() return an error: %v", err)
	}
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
	f, _ := Open(bookPath)
	defer f.Close()

	it, _ := f.Navigation()
	if it.Title() != firstTitle {
		t.Errorf("it.Title() return: %v when was expected: %v", it.Title(), firstTitle)
	}
}

func TestURL(t *testing.T) {
	f, _ := Open(bookPath)
	defer f.Close()

	it, _ := f.Navigation()
	if it.URL() != firstURL {
		t.Errorf("it.URL() return: %v when was expected: %v", it.URL(), firstURL)
	}
}

func TestDepth(t *testing.T) {
	f, _ := Open(bookPath)
	defer f.Close()

	it, _ := f.Navigation()
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
	if it.Title() != childTitle {
		t.Errorf("it.Title() return: %v when was expected: %v", it.Title(), childTitle)
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
	if it.Title() != firstTitle {
		t.Errorf("it.Title() return: %v when was expected: %v", it.Title(), firstTitle)
	}
}
