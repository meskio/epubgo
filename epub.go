package epub

import (
	"archive/zip"
)

type Epub struct {
	file *zip.ReadCloser
}

func Open(path string) (e Epub, err error) {
	e.file, err = zip.OpenReader(path)
	return e, err
}

func (e Epub) Close() {
	e.file.Close()
}
