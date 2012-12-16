// Copyright 2012 Ruben Pollan <meskio@sindominio.net>
// Use of this source code is governed by a LGPL licence
// version 3 or later that can be found in the LICENSE file.

package epub

import "testing"

const (
	book_path       = "testdata/a_dogs_tale.epub"
	book_title      = "A Dog's Tale"
	book_lang       = "en"
	book_identifier = "http://www.gutenberg.org/ebooks/3174"
	book_creator    = "Mark Twain"
	book_subject    = "Dogs -- Fiction"
	book_rights     = "Public domain in the USA."
)

func TestOpenClose(t *testing.T) {
	f, err := Open(book_path)
	if err != nil {
		t.Errorf("Open(%v) return an error: %v", book_path, err)
	}

	f.Close()
}

func TestMetadata(t *testing.T) {
	f, _ := Open(book_path)
	defer f.Close()

	if title, _ := f.Metadata("title"); title[0] != book_title {
		t.Errorf("Metadata title '%v', the expected was '%v'", title[0], book_title)
	}
	if language, _ := f.Metadata("language"); language[0] != book_lang {
		t.Errorf("Metadata language '%v', the expected was '%v'", language[0], book_lang)
	}
	if identifier, _ := f.Metadata("identifier"); identifier[0] != book_identifier {
		t.Errorf("Metadata identifier '%v', the expected was '%v'", identifier[0], book_identifier)
	}
	if creator, _ := f.Metadata("creator"); creator[0] != book_creator {
		t.Errorf("Metadata creator '%v', the expected was '%v'", creator[0], book_creator)
	}
	if subject, _ := f.Metadata("subject"); subject[0] != book_subject {
		t.Errorf("Metadata subject '%v', the expected was '%v'", subject[0], book_subject)
	}
	if rights, _ := f.Metadata("rights"); rights[0] != book_rights {
		t.Errorf("Metadata rights '%v', the expected was '%v'", rights[0], book_rights)
	}
}
