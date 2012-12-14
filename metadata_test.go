// Copyright 2012 Ruben Pollan <meskio@sindominio.net>
// Use of this source code is governed by a LGPL licence
// version 3 or later that can be found in the LICENSE file.

package epub

import "testing"

const (
	book_title      = "A Dog's Tale"
	book_lang       = "en"
	book_identifier = "http://www.gutenberg.org/ebooks/3174"
	book_creator    = "Mark Twain"
	book_subject    = "Dogs -- Fiction"
	book_rights     = "Public domain in the USA."
)

func TestMetadata(t *testing.T) {
	f, _ := Open(book_path)
	defer f.Close()

	if f.Metadata["title"][0].Content != book_title {
		t.Errorf("Metadata title '%v', the expected was '%v'", f.Metadata["title"], book_title)
	}
	if f.Metadata["language"][0].Content != book_lang {
		t.Errorf("Metadata language '%v', the expected was '%v'", f.Metadata["language"], book_lang)
	}
	if f.Metadata["identifier"][0].Content != book_identifier {
		t.Errorf("Metadata identifier '%v', the expected was '%v'", f.Metadata["identifier"], book_identifier)
	}
	if f.Metadata["creator"][0].Content != book_creator {
		t.Errorf("Metadata creator '%v', the expected was '%v'", f.Metadata["creator"], book_creator)
	}
	if f.Metadata["subject"][0].Content != book_subject {
		t.Errorf("Metadata subject '%v', the expected was '%v'", f.Metadata["subject"], book_subject)
	}
	if f.Metadata["rights"][0].Content != book_rights {
		t.Errorf("Metadata rights '%v', the expected was '%v'", f.Metadata["rights"], book_rights)
	}
}
