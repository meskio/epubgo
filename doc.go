// Copyright 2012 Ruben Pollan <meskio@sindominio.net>
// Use of this source code is governed by a LGPL licence
// version 3 or later that can be found in the LICENSE file.

/*
Package epubgo implements reading for epub ebook format.

A simple example of usage:
	book, err := epub.Open("path/of/the/file.epub")
	if err != nil {
		log.Panic(err)
	}
	defer book.Close()
	title, _ := book.Metadata("title")
	fmt.Println(title[0])

The pages of the book can be browsed with the SpineIterator:
	it, err := book.Spine()
	page := it.Open()
	defer page.Close()
	it.Next()

The index of the book can be browsed with the NavigationIterator:
	it, err := book.Navigation()
	it.Title()
	it.Next()
*/
package epubgo
