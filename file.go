// Copyright 2012 Ruben Pollan <meskio@sindominio.net>
// Use of this source code is governed by a LGPL licence
// version 3 or later that can be found in the LICENSE file.

package epubgo

import (
	"archive/zip"
	"encoding/xml"
	"errors"
	"golang.org/x/net/html/charset"
	"io"
	"path"
	"strings"
)

type containerXML struct {
	// FIXME: only support for one rootfile, can it be more than one?
	Rootfile rootfile `xml:"rootfiles>rootfile"`
}
type rootfile struct {
	Path string `xml:"full-path,attr"`
}

func openOPF(file *zip.Reader) (io.ReadCloser, error) {
	path, err := getOpfPath(file)
	if err != nil {
		return nil, err
	}
	return openFile(file, path)
}

func getRootPath(file *zip.Reader) (string, error) {
	opfPath, err := getOpfPath(file)
	if err != nil {
		return "", err
	}
	pathDir := path.Dir(opfPath)
	if pathDir == "." {
		return "", nil
	} else {
		return path.Dir(opfPath) + "/", nil
	}
}

func getOpfPath(file *zip.Reader) (string, error) {
	f, err := openFile(file, "META-INF/container.xml")
	if err != nil {
		return "", err
	}
	defer f.Close()

	var c containerXML
	err = decodeXML(f, &c)
	return c.Rootfile.Path, err
}

func decodeXML(file io.Reader, v interface{}) error {
	decoder := xml.NewDecoder(file)
	decoder.Entity = xml.HTMLEntity
	decoder.CharsetReader = charset.NewReaderLabel
	return decoder.Decode(v)
}

func openFile(file *zip.Reader, path string) (io.ReadCloser, error) {
	for _, f := range file.File {
		if f.Name == path {
			return f.Open()
		}
	}

	pathLower := strings.ToLower(path)
	for _, f := range file.File {
		if strings.ToLower(f.Name) == pathLower {
			return f.Open()
		}
	}

	return nil, errors.New("File " + path + " not found")
}
