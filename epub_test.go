package epub

import "testing"

const (
	book_path  = "testdata/a_dogs_tale.epub"
	book_title = "A Dog's Tale"
)

func TestOpenClose(t *testing.T) {
	f, err := Open(book_path)
	if err != nil {
		t.Errorf("Open(%v) return an error: %v", book_path, err)
	}

	f.Close()
}

func TestTitle(t *testing.T) {
	f, _ := Open(book_path)

	if f.Metadata.Title != book_title {
		t.Errorf("Metadata(title) return '%v', the expected was '%v'", f.Metadata.Title, book_title)
	}

	f.Close()
}
