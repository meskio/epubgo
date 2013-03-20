// Copyright 2012 Ruben Pollan <meskio@sindominio.net>
// Use of this source code is governed by a LGPL licence
// version 3 or later that can be found in the LICENSE file.

package epubgo

import (
	"archive/zip"
	"encoding/xml"
	"reflect"
	"strings"
)

// TODO convert date to date type?
type identifier struct {
	Data   string `xml:",chardata"`
	Id     string `xml:"id,attr"`
	Scheme string `xml:"scheme,attr"`
}
type author struct {
	Data   string `xml:",chardata"`
	FileAs string `xml:"file-as,attr"`
	Role   string `xml:"role,attr"`
}
type date struct {
	Data  string `xml:",chardata"`
	Event string `xml:"event,attr"`
}
type meta struct {
	Title       []string     `xml:"metadata>title"`
	Language    []string     `xml:"metadata>language"`
	Identifier  []identifier `xml:"metadata>identifier"`
	Creator     []author     `xml:"metadata>creator"`
	Subject     []string     `xml:"metadata>subject"`
	Description []string     `xml:"metadata>description"`
	Publisher   []string     `xml:"metadata>publisher"`
	Contributor []author     `xml:"metadata>contributor"`
	Date        []date       `xml:"metadata>date"`
	Type        []string     `xml:"metadata>type"`
	Format      []string     `xml:"metadata>format"`
	Source      []string     `xml:"metadata>source"`
	Relation    []string     `xml:"metadata>relation"`
	Coverage    []string     `xml:"metadata>coverage"`
	Rights      []string     `xml:"metadata>rights"`
}

func toMData(m meta) mdata {
	metadata := make(mdata)
	v := reflect.ValueOf(m)
	typeOf := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Len() == 0 {
			continue
		}

		fieldName := strings.ToLower(typeOf.Field(i).Name)
		data := make([]element, field.Len())
		for j := 0; j < field.Len(); j++ {
			data[j].attr = make(map[string]string)
			elem := field.Index(j).Interface()
			switch elem.(type) {
			case string:
				data[j].content, _ = elem.(string)
			case identifier:
				ident, _ := elem.(identifier)
				data[j].content = ident.Data
				data[j].attr["id"] = ident.Id
				data[j].attr["scheme"] = ident.Scheme
			case author:
				auth, _ := elem.(author)
				data[j].content = auth.Data
				data[j].attr["file-as"] = auth.FileAs
				data[j].attr["role"] = auth.Role
			case date:
				d, _ := elem.(date)
				data[j].content = d.Data
				data[j].attr["event"] = d.Event
			}
		}
		metadata[fieldName] = data
	}
	return metadata
}

func parseMetadata(file *zip.Reader) (metadata mdata, err error) {
	path, err := contentPath(file)
	if err != nil {
		return
	}

	f, err := openFile(file, path)
	if err != nil {
		return
	}
	defer f.Close()

	decoder := xml.NewDecoder(f)
	var m meta
	err = decoder.Decode(&m)
	if err != nil {
		return
	}

	metadata = toMData(m)
	return
}
