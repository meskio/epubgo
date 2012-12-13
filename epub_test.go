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

	if f.Metadata.Title[0] != book_title {
		t.Errorf("Metadata title '%v', the expected was '%v'", f.Metadata.Title, book_title)
	}
	if f.Metadata.Language[0] != book_lang {
		t.Errorf("Metadata language '%v', the expected was '%v'", f.Metadata.Language, book_lang)
	}
	if f.Metadata.Identifier[0] != book_identifier {
		t.Errorf("Metadata identifier '%v', the expected was '%v'", f.Metadata.Identifier, book_identifier)
	}
	if f.Metadata.Creator[0] != book_creator {
		t.Errorf("Metadata creator '%v', the expected was '%v'", f.Metadata.Creator, book_creator)
	}
	if f.Metadata.Subject[0] != book_subject {
		t.Errorf("Metadata subject '%v', the expected was '%v'", f.Metadata.Subject, book_subject)
	}
	if f.Metadata.Rights[0] != book_rights {
		t.Errorf("Metadata rights '%v', the expected was '%v'", f.Metadata.Rights, book_rights)
	}

	f.Close()
}
