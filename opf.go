// Copyright 2012 Ruben Pollan <meskio@sindominio.net>
// Use of this source code is governed by a LGPL licence
// version 3 or later that can be found in the LICENSE file.

package epubgo

import (
	"encoding/xml"
	"io"
	"reflect"
	"strings"
)

type xmlOPF struct {
	Metadata meta `xml:"metadata"`
}
type meta struct {
	Title       []string     `xml:"title"`
	Language    []string     `xml:"language"`
	Identifier  []identifier `xml:"identifier"`
	Creator     []author     `xml:"creator"`
	Subject     []string     `xml:"subject"`
	Description []string     `xml:"description"`
	Publisher   []string     `xml:"publisher"`
	Contributor []author     `xml:"contributor"`
	Date        []date       `xml:"date"`
	Type        []string     `xml:"type"`
	Format      []string     `xml:"format"`
	Source      []string     `xml:"source"`
	Relation    []string     `xml:"relation"`
	Coverage    []string     `xml:"coverage"`
	Rights      []string     `xml:"rights"`
}
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
	// TODO: convert date to date type?
	Data  string `xml:",chardata"`
	Event string `xml:"event,attr"`
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

func parseOPF(opf io.Reader) (metadata mdata, err error) {
	decoder := xml.NewDecoder(opf)
	var o xmlOPF
	err = decoder.Decode(&o)
	if err != nil {
		return
	}

	metadata = toMData(o.Metadata)
	return
}
