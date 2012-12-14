// Copyright 2012 Ruben Pollan <meskio@sindominio.net>
// Use of this source code is governed by a LGPL licence
// version 3 or later that can be found in the LICENSE file.

package epub

import (
	"archive/zip"
	"encoding/xml"
	"reflect"
	"strings"
)

type Data struct {
	Content string
	Attr    map[string]string
}
type MData map[string][]Data

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

func toMData(m meta) MData {
	metadata := make(MData)
	v := reflect.ValueOf(m)
	typeOf := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Len() == 0 {
			continue
		}

		fieldName := strings.ToLower(typeOf.Field(i).Name)
		data := make([]Data, field.Len())
		for j := 0; j < field.Len(); j++ {
			data[j].Attr = make(map[string]string)
			elem := field.Index(j).Interface()
			switch elem.(type) {
			case string:
				data[j].Content, _ = elem.(string)
			case identifier:
				ident, _ := elem.(identifier)
				data[j].Content = ident.Data
				data[j].Attr["id"] = ident.Id
				data[j].Attr["scheme"] = ident.Scheme
			case author:
				auth, _ := elem.(author)
				data[j].Content = auth.Data
				data[j].Attr["file-as"] = auth.FileAs
				data[j].Attr["role"] = auth.Role
			case date:
				d, _ := elem.(date)
				data[j].Content = d.Data
				data[j].Attr["event"] = d.Event
			}
		}
		metadata[fieldName] = data
	}
	return metadata
}

func parseMetadata(file *zip.ReadCloser) (metadata MData, err error) {
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
