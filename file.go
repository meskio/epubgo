package epub

import (
	"archive/zip"
	"encoding/xml"
	"errors"
	"io"
)

func openFile(file *zip.ReadCloser, path string) (io.ReadCloser, error) {
	for _, f := range file.File {
		if f.Name == path {
			return f.Open()
		}
	}
	return nil, errors.New("File " + path + " not found")
}

type rootfile struct {
	Path string `xml:"full-path,attr"`
}
type container_xml struct {
	// FIXME: only support for one rootfile, can it be more than one?
	Rootfile rootfile `xml:"rootfiles>rootfile"`
}

func contentPath(file *zip.ReadCloser) (string, error) {
	f, err := openFile(file, "META-INF/container.xml")
	if err != nil {
		return "", err
	}
	defer f.Close()

	var c container_xml
	decoder := xml.NewDecoder(f)
	err = decoder.Decode(&c)
	if err != nil {
		return "", err
	}
	return c.Rootfile.Path, nil
}
