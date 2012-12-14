package epub

import (
	"archive/zip"
)

type Data struct {
	Content string
	Attr    map[string]string
}
type MData map[string][]Data
type Epub struct {
	file     *zip.ReadCloser
	Metadata MData
}

func Open(path string) (e *Epub, err error) {
	e = new(Epub)
	e.file, err = zip.OpenReader(path)
	if err != nil {
		return
	}

	e.Metadata, err = parseMetadata(e.file)
	return
}

func (e Epub) Close() {
	e.file.Close()
}
