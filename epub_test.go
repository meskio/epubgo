package epub

import "testing"

const (
	book_path = "testdata/a_dogs_tale.epub"
)

func TestOpenClose(t *testing.T) {
	f, err := Open(book_path)
	if err != nil {
		t.Errorf("Open(%v) return an error: %v", book_path, err)
	}

	f.Close()
}
