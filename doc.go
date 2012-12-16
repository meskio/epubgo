// Copyright 2012 Ruben Pollan <meskio@sindominio.net>
// Use of this source code is governed by a LGPL licence
// version 3 or later that can be found in the LICENSE file.

/*
Package epub implements reading for epub ebook format.

A simple example of usage:
	book, err := epub.Open("path/of/the/file.epub")
	if err != nil {
		log.Panic(err)
	}
	difer book.Close()
	title, _ := book.Metadata("title")
	fmt.Println(title[0])

*/
package epub
