// Copyright 2012 Ruben Pollan <meskio@sindominio.net>
// Use of this source code is governed by a LGPL licence
// version 3 or later that can be found in the LICENSE file.

package epubgo

import _ "code.google.com/p/go-charset/data"

import (
	"archive/zip"
	"code.google.com/p/go-charset/charset"
	"encoding/xml"
	"errors"
	"io"
	"strings"
)

type container_xml struct {
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
	index := strings.LastIndex(opfPath, "/")
	if index == -1 {
		return "", nil
	}
	return opfPath[:index+1], nil
}

func getOpfPath(file *zip.Reader) (string, error) {
	f, err := openFile(file, "META-INF/container.xml")
	if err != nil {
		return "", err
	}
	defer f.Close()

	var c container_xml
	err = decodeXml(f, &c)
	return c.Rootfile.Path, err
}

func decodeXml(file io.Reader, v interface{}) error {
	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReader
	return decoder.Decode(v)
}

func openFile(file *zip.Reader, path string) (io.ReadCloser, error) {
	for _, f := range file.File {
		if f.Name == path {
			return f.Open()
		}
	}
	return nil, errors.New("File " + path + " not found")
}
