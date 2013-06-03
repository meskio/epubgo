// Copyright 2012 Ruben Pollan <meskio@sindominio.net>
// Use of this source code is governed by a LGPL licence
// version 3 or later that can be found in the LICENSE file.

package epubgo

import "testing"

import (
	"bytes"
	"io/ioutil"
)

const (
	spineURL = "wrap0000.html"
)

func TestFirst(t *testing.T) {
	f, _ := Open(bookPath)
	defer f.Close()

	it, err := f.Spine()
	if err != nil {
		t.Errorf("epub.Spine() return an error: %v", err)
	}
	if !it.IsFirst() {
		t.Errorf("it.IsFirst() not behaving as expected")
	}
	if it.IsLast() {
		t.Errorf("it.IsLast() not behaving as expected")
	}
}

func TestLast(t *testing.T) {
	f, _ := Open(bookPath)
	defer f.Close()

	it, _ := f.Spine()
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

func TestLastNext(t *testing.T) {
	f, _ := Open(bookPath)
	defer f.Close()

	it, _ := f.Spine()
	if err := it.Next(); err != nil {
		t.Errorf("it.Next() return an error: %v", err)
	}
	if err := it.Next(); err == nil {
		t.Errorf("it.Next() didn't return an error being the last element")
	}
}

func TestMove(t *testing.T) {
	f, _ := Open(bookPath)
	defer f.Close()

	it, _ := f.Spine()
	if it.Previous() == nil {
		t.Errorf("it.Previous() din't return an error being the first")
	}
	if err := it.Next(); err != nil {
		t.Errorf("it.Next() return an error: %v", err)
	}
	if err := it.Next(); err == nil {
		t.Errorf("it.Next() didn't return an error being the last")
	}
	if err := it.Previous(); err != nil {
		t.Errorf("it.Next() return an error: %v", err)
	}
	if !it.IsFirst() {
		t.Errorf("it.IsFirst() not behaving as expected")
	}
	if it.IsLast() {
		t.Errorf("it.IsLast() not behaving as expected")
	}
}

func TestSpineURL(t *testing.T) {
	f, _ := Open(bookPath)
	defer f.Close()

	it, _ := f.Spine()
	if it.URL() != spineURL {
		t.Errorf("it.URL() return: %v when was expected: %v", it.URL(), spineURL)
	}
}

func TestSpineOpen(t *testing.T) {
	f, _ := Open(bookPath)
	defer f.Close()

	it, _ := f.Spine()
	html1, err := it.Open()
	if err != nil {
		t.Errorf("it.Open() return an error: %v", err)
		return
	}
	defer html1.Close()

	html2, err := f.OpenFile(it.URL())
	if err != nil {
		t.Errorf("OpenFile(%v) return an error: %v", it.URL(), err)
		return
	}
	defer html2.Close()

	buff1, err := ioutil.ReadAll(html1)
	if err != nil {
		t.Errorf("Error reading the opened file: %v", err)
		return
	}
	buff2, _ := ioutil.ReadAll(html2)
	if !bytes.Equal(buff1, buff2) {
		t.Errorf("The files on epub and spine iterator are not equal")
		return
	}
}
